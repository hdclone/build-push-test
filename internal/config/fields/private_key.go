package fields

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"
)

type PrivateKey struct {
	Key     *ecdsa.PrivateKey
	Address common.Address
}

func (pk *PrivateKey) Decode(value string) error {
	return pk.parse(value)
}

func (pk *PrivateKey) UnmarshalYAML(value *yaml.Node) error {
	return pk.parse(value.Value)
}

func (pk *PrivateKey) parse(value string) error {
	if key, err := ethCrypto.HexToECDSA(value); err != nil {
		return err
	} else {
		pk.Key = key
	}
	pk.Address = ethCrypto.PubkeyToAddress(*pk.Key.Public().(*ecdsa.PublicKey))
	return nil
}
