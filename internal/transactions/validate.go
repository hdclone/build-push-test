package transactions

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func ValidateSignature(mpc []byte, callData []byte, signature []byte, receiveSide common.Address, chainID int64, bridgeAddress common.Address) error {
	var (
		err           error
		publicKey     *ecdsa.PublicKey
		publicAddress common.Address
	)

	//// https://www.google.com/search?q=ecrecover+27+28
	//if signature[len(signature)-1] <= 1 {
	//	signature[len(signature)-1] += 27
	//}

	var data []byte
	data = append(data, []byte("receiveRequestV2")...)
	data = append(data, callData...)
	data = append(data, receiveSide.Bytes()...)
	data = append(data, math.U256Bytes(big.NewInt(chainID))...)
	data = append(data, bridgeAddress.Bytes()...)

	hash := crypto.Keccak256Hash(data)

	publicKey, err = crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return err
	}

	publicAddress = crypto.PubkeyToAddress(*publicKey)

	if !bytes.Equal(mpc, publicAddress.Bytes()) {
		return fmt.Errorf(
			"recovered address (%s) does not match current MPC address (%s)",
			publicAddress.String(),
			common.BytesToAddress(mpc).String(),
		)
	}

	return nil
}
