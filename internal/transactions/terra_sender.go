package transactions

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/client"
	tx2 "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
	"github.com/terra-money/core/app"
	core "github.com/terra-money/core/types"
	wasmtypes "github.com/terra-money/core/x/wasm/types"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"broadcaster/internal/config"
	"broadcaster/internal/logging"
	"broadcaster/internal/model"
	"broadcaster/internal/modules"
	repository "broadcaster/internal/repository/transactions"
)

type ReceiveRequest struct {
	CallData    []byte `json:"call_data"`
	Signature   []byte `json:"signature"`
	ReceiveSide []byte `json:"receive_side"`
}

type ContractRequest struct {
	Method ReceiveRequest `json:"receive_request"`
}

type StdFee struct {
	Amount sdk.Coins
	Gas    uint64
}

func TerraSend(
	ctx context.Context,
	transaction *model.Transaction,
	chainConfig *config.ChainConfig,
	logger *zap.Logger,
) (retryErr error) {
	var (
		gasLimit uint64
		GasPrice *big.Int
	)

	grpcConn := modules.TerraConn(chainConfig)
	txClient := tx.NewServiceClient(grpcConn)

	// request sas price from advisor
	advisorResult, err := requestGasAdvisor(ctx, transaction, logger)
	if err != nil {
		return err
	}

	pk := &secp256k1.PrivKey{
		Key: chainConfig.Key.Key.D.Bytes(),
	}
	encCfg := app.MakeEncodingConfig()

	// Create a new TxBuilder.
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	addrFrom := sdk.AccAddress(pk.PubKey().Address())

	authClient := authtypes.NewQueryClient(grpcConn)
	encCfg.InterfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	accountRes, err := authClient.Account(ctx,
		&authtypes.QueryAccountRequest{
			Address: addrFrom.String(),
		},
	)
	if err != nil {
		return err
	}
	cdk := codec.NewProtoCodec(encCfg.InterfaceRegistry)
	var account authtypes.AccountI
	err = cdk.UnpackAny(accountRes.Account, &account)
	if err != nil {
		return err
	}

	if advisorResult.GasPriceE1559 {
		GasPrice = nil
	}

	// set transaction status to sending
	addrContract := transaction.BridgeAddress.Bytes()

	contrMethod := ContractRequest{Method: ReceiveRequest{
		CallData:    transaction.CallData,
		Signature:   transaction.Signature,
		ReceiveSide: transaction.ReceiveSide.Bytes(),
	}}

	contrData, err := json.Marshal(contrMethod)
	if err != nil {
		return err
	}

	msgContract := wasmtypes.NewMsgExecuteContract(addrFrom, addrContract, contrData, nil)
	err = txBuilder.SetMsgs(msgContract)
	if err != nil {
		return err
	}

	sdtFees, err := calcGas(ctx, chainConfig.Terra.ID, encCfg.TxConfig, account.GetSequence(), account.GetAccountNumber(), grpcConn, msgContract)
	if err != nil {
		return err
	}
	if chainConfig.GasLimit > 0 {
		gasLimit = chainConfig.GasLimit
	} else {
		if sdtFees.Gas < 1000000 {
			gasLimit = 1000000
			logger.Warn(fmt.Sprintf("estimated gas %d was changed to %d", sdtFees.Gas, gasLimit))
		} else {
			gasLimit = sdtFees.Gas * 120 / 100
		}
	}

	sendingStatusErr := repository.SetSendingStatus(ctx, transaction, account.GetSequence(), gasLimit, GasPrice)

	if sendingStatusErr != nil {
		return errors.WithMessage(sendingStatusErr, "can't set sending status for transaction")
	}

	txBuilder.SetGasLimit(gasLimit)
	txBuilder.SetFeeAmount(sdtFees.Amount)
	//TODO research what are these options
	txBuilder.SetFeeGranter(addrFrom)

	sigV2 := signing.SignatureV2{
		PubKey: pk.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	sigV2 = signing.SignatureV2{}
	signerData := xauthsigning.SignerData{
		ChainID:       chainConfig.Terra.ID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}
	sigV2, err = tx2.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, pk, encCfg.TxConfig, account.GetSequence())
	if err != nil {
		return err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}

	broadcastRes, err := txClient.BroadcastTx(
		ctx,
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return err
	}
	fmt.Println(broadcastRes)

	trxHash := []byte(broadcastRes.TxResponse.TxHash)

	return repository.SetSentStatus(logging.CtxSet(ctx, logger.With(zap.String("tx_hash", broadcastRes.TxResponse.TxHash))), transaction, trxHash)
}

//calcGas return GasWanted
func calcGas(ctx context.Context, chainId string, txConfig client.TxConfig, sequence uint64, accountNumber uint64, grpcConn *grpc.ClientConn, msgContract sdk.Msg) (*StdFee, error) {
	txf := tx2.Factory{}.
		WithChainID(chainId).
		WithTxConfig(txConfig).
		WithSequence(sequence).
		WithAccountNumber(accountNumber).
		WithGasPrices("1" + core.MicroLunaDenom)

	txBytes, err := tx2.BuildSimTx(txf, msgContract)
	if err != nil {
		return nil, err
	}
	txSvcClient := tx.NewServiceClient(grpcConn)

	simRes, err := txSvcClient.Simulate(ctx, &tx.SimulateRequest{
		TxBytes: txBytes,
	})
	if err != nil {
		return nil, err
	}
	gas := simRes.GasInfo.GasWanted

	fees := txf.Fees()
	gasPrices := txf.GasPrices()

	if !gasPrices.IsZero() {
		glDec := sdk.NewDec(int64(gas))

		// Derive the fees based on the provided gas prices, where
		// fee = ceil(gasPrice * gasLimit).
		gasFees := make(sdk.Coins, len(gasPrices))
		for i, gp := range gasPrices {
			fee := gp.Amount.Mul(glDec)
			gasFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}

		fees = fees.Add(gasFees.Sort()...)
	}

	return &StdFee{
		Amount: fees,
		Gas:    gas,
	}, nil
}
