package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	wasmtypes "github.com/terra-money/core/x/wasm/types"
	"go.uber.org/zap"

	"broadcaster/internal/bridge"
	"broadcaster/internal/config"
	"broadcaster/internal/logging"
	"broadcaster/internal/model"
	"broadcaster/internal/modules"
	"broadcaster/internal/pool"
	repository "broadcaster/internal/repository/transactions"
	"broadcaster/internal/response"
	"broadcaster/internal/transactions"
	"broadcaster/internal/validator"
	"broadcaster/internal/validator/rules"
)

type RequestBody struct {
	CallData    []byte `json:"call_data"`
	Signature   []byte `json:"signature"`
	ReceiveSide string `json:"receive_side"`
	TxHash      string `json:"tx_hash"`
}

type RequestBroadcast struct {
	ChainIDFrom   int64          `json:"chain_id_from"`
	ChainID       int64          `json:"chain_id"`
	BridgeAddress common.Address `json:"bridge_address"`
	RequestBody
}

type TerraContractRes struct {
	Mpc string `json:"mpc"`
}

func Broadcast(w http.ResponseWriter, r *http.Request) {
	var (
		err               error
		ctx               = r.Context()
		params            = mux.Vars(r)
		requestBroadcast  = RequestBroadcast{}
		validationResults *validator.Results
		logger            = logging.CtxGet(r.Context())
	)

	logger.Debug("parse body")
	if err = json.NewDecoder(r.Body).Decode(&requestBroadcast.RequestBody); err != nil {
		panic(response.BadRequest().WithMessage("unable to parse request body").WithError(errors.WithStack(err)))
	}

	requestBroadcast.ChainIDFrom, _ = strconv.ParseInt(params["chainIdFrom"], 10, 32)
	requestBroadcast.ChainID, _ = strconv.ParseInt(params["chainId"], 10, 32)

	receiveSide := common.HexToAddress(requestBroadcast.ReceiveSide)
	chainConfig, err := config.CtcGet(ctx).Chains.Get(requestBroadcast.ChainID)
	if err != nil {
		return
	}

	switch chainConfig.Kind {
	case config.KindEVM:
		requestBroadcast.BridgeAddress = common.HexToAddress(params["bridge"])
	case config.KindCosmos:
		bridgeAddress, err := sdk.GetFromBech32(params["bridge"], "terra")
		if err != nil {
			logger.With(zap.Error(err), zap.String("bridge_from_url", params["bridge"])).Error("failed to convert Cosmos bridge from Bech32")
			return
		}
		requestBroadcast.BridgeAddress = common.BytesToAddress(bridgeAddress)
		logger.With(
			zap.Error(err),
			zap.String("bridge_address_url", params["bridge"]),
			zap.Int("bytes_length", len(bridgeAddress)),
			zap.String("bridge_address_result", requestBroadcast.BridgeAddress.String()),
		).Debug("converted Cosmos bridge address to ETH address from Bech32")
	default:
		return
	}

	txHash := common.HexToHash(requestBroadcast.TxHash)

	logger = logger.With(zap.String("tx_hash_source", requestBroadcast.TxHash))

	logger.Info("validate request")
	if validationResults, err = validator.Scope(
		validator.Validate("chainIdFrom", requestBroadcast.ChainIDFrom).Rules(rules.IsCustom(func() error {
			if _, err = modules.Config().Chains.Get(requestBroadcast.ChainIDFrom); err != nil {
				return validator.Message(err.Error())
			}
			return nil
		})),
		validator.Validate("chainId", requestBroadcast.ChainID).Rules(rules.IsCustom(func() error {
			if _, err = modules.Config().Chains.Get(requestBroadcast.ChainID); err != nil {
				return validator.Message(err.Error())
			}
			return nil
		})),
		validator.Validate("bridge", requestBroadcast.BridgeAddress).Rules(rules.IsCustom(func() error {
			if chainConfig, err := config.CtcGet(ctx).Chains.Get(requestBroadcast.ChainID); err != nil {
				return validator.Message("can't found chain of bridge")
			} else if !chainConfig.Whitelist.Accepted(requestBroadcast.BridgeAddress) {
				return validator.Message("bridge not in whitelist")
			}
			return nil
		})),
		validator.Validate("signature", requestBroadcast.Signature).Rules(rules.IsPresence()),
		validator.Validate("call_data", requestBroadcast.CallData).Rules(rules.IsPresence()),
		validator.Validate("receive_side", receiveSide.Bytes()).Rules(rules.IsPresence()),
		validator.Validate("tx_hash", requestBroadcast.TxHash).Rules(rules.IsPresence()),
	).Validate(ctx); err != nil {
		panic(response.InternalServerError().WithError(err))
	} else if !validationResults.Valid() {
		logger.With(
			zap.Any("validation", validationResults),
		).Warn("invalid request")
		panic(response.UnprocessableEntity().WithMessage("invalid request").WithData(validationResults))
	}

	ctx = logging.CtxSet(ctx, logger)

	logger.Info("get MPC address from bridge")
	mpcAddr, err := mpc(ctx, chainConfig, requestBroadcast.BridgeAddress)
	if err != nil {
		logger.With(zap.Error(err)).Warn("err get mpc")
		panic(response.BadRequest().WithMessage("err get mpc").WithError(errors.WithStack(err)))
	}

	logger.With(
		zap.String("mpc_from_bridge", common.BytesToAddress(mpcAddr).String()),
	).Info("validate signature")

	if err = transactions.ValidateSignature(mpcAddr, requestBroadcast.CallData, requestBroadcast.Signature, receiveSide, requestBroadcast.ChainID, requestBroadcast.BridgeAddress); err != nil {
		logger.With(zap.Error(err)).Warn("bad signature")
		panic(response.BadRequest().WithMessage("bad signature").WithError(errors.WithStack(err)))
	}

	logger.Info("check for unique transaction")
	var transactionPresent bool
	if transactionPresent, err = repository.Query(ctx).Column("*").Apply(repository.ByRequest(requestBroadcast.ChainID, requestBroadcast.BridgeAddress, receiveSide, requestBroadcast.CallData)).Exists(ctx); err != nil {
		panic(response.InternalServerError().WithError(err))
	} else if transactionPresent {
		logger.Warn("transaction is not unique")
		panic(response.BadRequest().WithMessage("transaction is not unique"))
	}

	// Store transaction as pending
	logger.Info("store pending transaction")
	var transaction *model.Transaction
	transaction, err = repository.Create(
		ctx,
		requestBroadcast.ChainIDFrom,
		requestBroadcast.ChainID,
		requestBroadcast.BridgeAddress,
		receiveSide,
		requestBroadcast.CallData,
		requestBroadcast.Signature,
		txHash,
		chainConfig.Key.Address,
		model.NetworkType(chainConfig.Kind),
	)
	if err != nil {
		panic(response.InternalServerError().WithError(errors.WithStack(err)))
	}

	if err = modules.Queue(transaction.ChainID).Enqueue(transaction); err != nil {
		panic(response.InternalServerError().WithMessage("can't enqueue transaction").WithError(err))
	}

	response.Ok().WithMessage("transaction was sent").Write(w, r)
}

func mpc(ctx context.Context, chainConfig *config.ChainConfig, bridgeAddress common.Address) (mpc []byte, err error) {
	var logger = logging.CtxGet(ctx)
	logger.Debug(fmt.Sprintf("Getting MPC address from bridge of kind %s", chainConfig.Kind))
	switch chainConfig.Kind {
	case config.KindEVM:
		poolByChain := pool.CtcGet(ctx).Clients(chainConfig)
		err = poolByChain.Try(ctx, func(ctx context.Context, client *ethclient.Client) error {
			var (
				rpcErr   error
				contract *bridge.Bridge
			)
			contract, rpcErr = bridge.NewBridge(bridgeAddress, client)
			if rpcErr != nil {
				return errors.WithStack(rpcErr)
			}
			evmMpc, err := contract.Mpc(nil)
			if err != nil {
				return err
			}
			mpc = evmMpc.Bytes()

			return nil
		})
	case config.KindCosmos:
		client := modules.TerraConn(chainConfig)
		wClient := wasmtypes.NewQueryClient(client)
		var contractAddr sdk.AccAddress
		contractAddr = bridgeAddress.Bytes()
		daraR := []byte("{\"get_mpc\": {}}")
		info, err := wClient.ContractStore(ctx, &wasmtypes.QueryContractStoreRequest{
			ContractAddress: contractAddr.String(),
			QueryMsg:        daraR,
		})
		if err != nil {
			logger.With(zap.Error(err)).Warn("Failed to create contract store while getting MPC")
			break
		}
		res := TerraContractRes{}
		err = json.Unmarshal(info.GetQueryResult(), &res)
		if err != nil {
			logger.With(zap.Error(err)).Warn("Failed to unmarshal query result for get_mpc on terra")
			break
		}
		mpcAddr, err := sdk.GetFromBech32(res.Mpc, "terra")
		if err != nil {
			logger.With(
				zap.Error(err),
				zap.String("mpc_cosmos", res.Mpc),
			).Warn("Failed to convert Bech32 MPC address to account address")
			break
		} else {
			logger.With(
				zap.String("mpc_cosmos", res.Mpc),
				zap.String("mpc_eth", common.BytesToAddress(mpcAddr).String()),
			).Debug("Converted Bech32 MPC address to account address")
		}

		mpc = mpcAddr
	default:
		err = fmt.Errorf("wrong chain type, received: %s", chainConfig.Kind)
	}
	return mpc, err
}
