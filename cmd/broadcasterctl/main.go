package main

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"broadcaster/internal/config"
	"broadcaster/internal/database"
	"broadcaster/internal/model"
	"broadcaster/internal/modules"
	repository "broadcaster/internal/repository/transactions"
	"broadcaster/internal/transactions"
	"broadcaster/internal/variables"
)

func main() {
	modules.Init()

	logger := modules.Logger()
	defer func() {
		_ = logger.Sync()
	}()

	app := &cli.App{
		Name:     variables.Banner("Broadcaster CTL"),
		Version:  variables.Version,
		Usage:    "tools of broadcaster service",
		HideHelp: false,
		Commands: []*cli.Command{{
			Name:        "transactions",
			Description: "manage transactions",
			Subcommands: []*cli.Command{
				{
					Name:        "all",
					Description: "list all transactions",
					Action: func(c *cli.Context) error {
						return ListTransactions()
					},
				},
				{
					Name:  "pending",
					Usage: "list pending transactions",
					Action: func(c *cli.Context) error {
						return ListTransactions(model.TransactionStatusPending)
					},
				},
				{
					Name:  "failed",
					Usage: "list failed transactions",
					Action: func(c *cli.Context) error {
						return ListTransactions(model.TransactionStatusFailed)
					},
				},
				{
					Name:  "sent",
					Usage: "list failed transactions",
					Action: func(c *cli.Context) error {
						return ListTransactions(model.TransactionStatusSent)
					},
				},
				{
					Name:  "resend",
					Usage: "resend transaction",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name: "id",
						},
					},
					Action: func(c *cli.Context) error {
						id := c.String("id")
						if len(id) == 0 {
							return errors.New("id must be passed")
						}
						return ResendTransaction(model.ID(id))
					},
				},
			},
		}},
	}
	if err := app.Run(os.Args); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func ResendTransaction(transactionId model.ID) error {
	ctx := database.CtxSet(context.Background(), modules.Database())
	ctx = config.CtxSet(ctx, modules.Config())
	var transaction model.Transaction
	if _, err := repository.Query(ctx).Apply(repository.ID(transactionId)).Exec(ctx, &transaction); errors.Is(err, sql.ErrNoRows) {
		return errors.Errorf("transaction with id '%s' not found", transactionId)
	} else if err != nil {
		return err
	}
	return transactions.Sender(ctx, &transaction, false)
}

func ListTransactions(status ...model.TransactionStatus) error {
	ctx := database.CtxSet(context.Background(), modules.Database())
	var transactionsByStatus []*model.Transaction
	query := repository.Query(ctx)
	if len(status) > 0 {
		query = query.Apply(repository.Status(status...))
	}
	if _, err := query.Exec(ctx, &transactionsByStatus); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"Status",
		"Error",
		"Bridge",
		"To Chain",
		"From Chain",
		"Receive size",
		"Gas Limit",
		"Gas price",
		"Created At",
		"Updated At",
	})
	for _, transaction := range transactionsByStatus {

		gasLimit := ""
		gasPrice := ""
		trxError := ""
		createdAt := ""
		updatedAt := ""

		if transaction.GasLimit != nil {
			gasLimit = strconv.FormatUint(*transaction.GasLimit, 10)
		}
		if transaction.GasPrice != nil {
			value := big.Int(*transaction.GasPrice)
			gasPrice = fmt.Sprintf("%v", value.String())
		}

		if transaction.Error != nil {
			trxError = *transaction.Error
		}

		if transaction.UpdatedAt != nil {
			updatedAt = transaction.UpdatedAt.Format("02 Jan 2006 15:04:05 MST")
		}

		if transaction.CreatedAt != nil {
			createdAt = transaction.CreatedAt.Format("02 Jan 2006 15:04:05 MST")
		}

		table.Append([]string{
			string(transaction.ID),
			string(transaction.Status),
			trxError,
			transaction.BridgeAddress.String(),
			strconv.FormatInt(transaction.ChainID, 10),
			strconv.FormatInt(transaction.ChainIDFrom, 10),
			transaction.ReceiveSide.String(),
			gasLimit,
			gasPrice,
			createdAt,
			updatedAt,
		})
	}
	table.Render()
	return nil
}
