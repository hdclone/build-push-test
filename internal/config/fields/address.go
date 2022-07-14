package fields

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gopkg.in/yaml.v3"
)

type Address common.Address

func (a *Address) parseValue(rawValue string) error {
	return hexutil.UnmarshalFixedText("Address", []byte(rawValue), a[:])
}

func (a *Address) UnmarshalYAML(value *yaml.Node) error {
	return a.parseValue(value.Value)
}

func (a *Address) UnmarshalJSON(value []byte) error {
	return a.parseValue(string(value))
}

func (a *Address) Decode(rawValue string) error {
	return a.parseValue(rawValue)
}

func (a *Address) ToString() string {
	return common.Address(*a).String()
}

func (a Address) ToCommon() common.Address {
	return common.Address(a)
}
