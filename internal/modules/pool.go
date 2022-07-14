package modules

import (
	"broadcaster/internal/pool"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Pool() *pool.Pool {
	return Register("pool", func(s string) (Module, error) {
		return pool.New(func(rpc string) (*ethclient.Client, error) {
			return ethclient.Dial(rpc)
		}), nil
	}).(*pool.Pool)
}
