package model

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

type BigInt big.Int

func (container *BigInt) MarshalJSON() ([]byte, error) {
	return []byte((*big.Int)(container).String()), nil
}

func (container *BigInt) Value() (driver.Value, error) {
	if container != nil {
		return (*big.Int)(container).String(), nil
	}
	return nil, nil
}

func (container *BigInt) Scan(value interface{}) error {
	if value == nil {
		container = nil
	}

	switch value := value.(type) {
	case []byte:
		return container.Scan(string(value))
	case string:
		_, ok := (*big.Int)(container).SetString(value, 10)
		if !ok {
			return fmt.Errorf("failed to load value to []byte: %v", value)
		}
	default:
		return fmt.Errorf("could not scan type %T into BigInt", value)
	}

	return nil
}
