package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	core "github.com/terra-money/core/types"

	"broadcaster/cmd/broadcasterd/api"
	v1 "broadcaster/cmd/broadcasterd/api/v1"
	"broadcaster/internal/config"
	"broadcaster/internal/database"
	"broadcaster/internal/handlers"
	"broadcaster/internal/logging"
	"broadcaster/internal/middlewares"
	"broadcaster/internal/model"
	"broadcaster/internal/modules"
	repository "broadcaster/internal/repository/transactions"
	"broadcaster/internal/transactions"
	"broadcaster/internal/variables"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/pkger"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func init() {
	goqu.SetDefaultPrepared(true)
}

func main() {
	modules.Init()
	cfg := modules.Config()
	logger := modules.Logger()

	sdk.GetConfig().SetBech32PrefixForAccount(core.Bech32PrefixAccAddr, core.Bech32PrefixAccPub)
	sdk.GetConfig().SetCoinType(core.CoinType)
	logger.Info(variables.Banner("Broadcaster Service"))

	defer func() {
		exit := false
		if err := recover(); err != nil {
			logger.Error(fmt.Sprintf("panic := %s", err))
			logger.Error("stacktrace from panic: \n" + string(debug.Stack()))
			exit = true
		}

		logger.Info("shutdown")
		_ = logger.Sync()
		if exit {
			os.Exit(1)
		}
	}()

	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.Initialize, middlewares.Catch, middlewares.Logger, middlewares.CORS)

	handlers.Default(router)

	api.Mount(router, "", v1.Router().StrictSlash(true))

	httpServer := &http.Server{
		Addr:    cfg.Server.Address.ToString(),
		Handler: router,
	}

	ctx := database.CtxSet(context.Background(), modules.Database())
	ctx = logging.CtxSet(ctx, logger)
	ctx = config.CtxSet(ctx, modules.Config())

	for _, chain := range modules.Config().Chains {
		modules.Queue(chain.ID).Start(ctx, chain, func(ctx context.Context, transaction *model.Transaction) error {
			return transactions.Sender(ctx, transaction, true)
		})
	}

	if cfg.Recovery.Enabled {
		var pendingTransactions []*model.Transaction
		if _, err := repository.Query(ctx).Apply(repository.Status(model.TransactionStatusPending, model.TransactionStatusSending)).
			Exec(ctx, &pendingTransactions); err != nil && !errors.Is(err, sql.ErrNoRows) {
			logger.Error(err.Error())
		} else {
			for _, transaction := range pendingTransactions {
				if err = modules.Queue(transaction.ChainID).Enqueue(transaction); err != nil {
					logger.Error(err.Error())
				}
			}
		}
	}

	go func() {
		logger.Info(fmt.Sprintf("listen http://%s", cfg.Server.Address.ToString()))
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(len(modules.Config().Chains))
	for _, chain := range modules.Config().Chains {
		go func(wg *sync.WaitGroup, c *config.ChainConfig) {
			if err := modules.Queue(c.ID).Shutdown(); err != nil {
				logger.Error(err.Error())
			}
			wg.Done()
		}(&wg, chain)
	}
	wg.Wait()
}
