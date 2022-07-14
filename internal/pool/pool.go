package pool

import (
	"broadcaster/internal/config"
	"broadcaster/internal/database"
	"broadcaster/internal/logging"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type Connector func(rpc string) (*ethclient.Client, error)
type Callback func(ctx context.Context, client *ethclient.Client) error

type Pool struct {
	mutex     sync.Mutex
	clients   map[string]*ethclient.Client
	connector Connector
}

func New(connector Connector) *Pool {
	return &Pool{
		clients:   map[string]*ethclient.Client{},
		mutex:     sync.Mutex{},
		connector: connector,
	}
}

func (ec *Pool) newClient(clientId string, endpoint string) (*ethclient.Client, error) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	newClient, err := ec.connector(endpoint)
	if err != nil {
		return nil, err
	}
	ec.clients[clientId] = newClient
	return newClient, nil
}

type Clients struct {
	pool        *Pool
	chainConfig *config.ChainConfig
}

func (ec *Pool) Clients(chainConfig *config.ChainConfig) *Clients {
	return &Clients{
		pool:        ec,
		chainConfig: chainConfig,
	}
}

func (ec *Clients) Try(ctx context.Context, callback Callback) error {
	var wg sync.WaitGroup
	result := make(chan bool, len(ec.chainConfig.Endpoints))
	c, cancel := context.WithTimeout(ctx, time.Minute*5)

	for rpcIndex := 0; rpcIndex < len(ec.chainConfig.Endpoints); rpcIndex++ {
		clientId := fmt.Sprintf("%v-%v", ec.chainConfig.ID, rpcIndex)
		client, foundClient := ec.pool.clients[clientId]
		endpoint := ec.chainConfig.Endpoints[rpcIndex]
		callLogger := logging.CtxGet(c).With(zap.String("endpoint", endpoint), zap.Int64("chainId", ec.chainConfig.ID))
		if !foundClient {
			var err error
			client, err = ec.pool.newClient(clientId, endpoint)
			if err != nil {
				continue
			}
		}

		wg.Add(1)
		go func(c context.Context, client *ethclient.Client, endpoint string, logger *zap.Logger) {
			defer wg.Done()

			logger.Debug("rpc request")
			if err := callback(logging.CtxSet(c, logger), client); err != nil {
				logger.Error(err.Error())
				logger.Error("rpc request error")
				result <- false
				return
			}

			result <- true
			logger.Debug("rpc request success")
		}(database.CtxSet(c, database.CtxGet(ctx)), client, endpoint, callLogger)
	}

	go func() {
		wg.Wait()
		cancel()
		close(result)
	}()

	for ok := range result {
		if ok {
			return nil
		}
	}

	return fmt.Errorf("all rpc urls returned error")
}
