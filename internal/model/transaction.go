package model

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	eth_common "github.com/ethereum/go-ethereum/common"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusSending   TransactionStatus = "sending"
	TransactionStatusSent      TransactionStatus = "sent"
	TransactionStatusConfirmed TransactionStatus = "confirmed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

type NetworkType string

const (
	ChainTypeEvm    NetworkType = "evm"
	ChainTypeCosmos NetworkType = "cosmos"
)

const (
	TransactionsTable = "transactions"
	TransactionsKind  = "trx"
)

type Transaction struct {
	ID                       ID                 `json:"id" db:"id" bun:",pk"`
	CreatedAt                *time.Time         `json:"created_at,omitempty" db:"created_at" goqu:"skipinsert"`
	ChainIDFrom              int64              `json:"chain_id_from" db:"chain_id_from"`
	ChainID                  int64              `json:"chain_id" db:"chain_id"`
	BridgeAddress            eth_common.Address `json:"bridge_address" db:"bridge_address"`
	ReceiveSide              eth_common.Address `json:"receive_side" db:"receive_side"`
	CallData                 []byte             `json:"call_data" db:"call_data"`
	Signature                []byte             `json:"signature" db:"signature"`
	Status                   TransactionStatus  `json:"status,omitempty" db:"status"`
	SenderAddress            eth_common.Address `json:"sender_address" db:"sender_address"`
	UpdatedAt                *time.Time         `json:"updated_at,omitempty" db:"updated_at" goqu:"skipinsert"`
	Nonce                    *uint64            `json:"nonce" db:"nonce"`
	GasLimit                 *uint64            `json:"gas_limit,omitempty" db:"gas_limit"`
	GasPrice                 *BigInt            `json:"gas_price,omitempty" db:"gas_price"`
	Hash                     []byte             `json:"hash,omitempty" db:"hash"`
	HashSource               []byte             `json:"hash_source,omitempty" db:"hash_source"`
	Error                    *string            `json:"error,omitempty" db:"error"`
	SourceTransactionMinedAt *time.Time         `json:"source_transaction_mined_at,omitempty" db:"source_transaction_mined_at"`
	NetworkType              *NetworkType       `json:"network_type,omitempty" db:"chain_type"`
}

func (transaction *Transaction) Sending(nonce uint64, gasLimit uint64, gasPrice *BigInt) {
	transaction.Nonce = &nonce
	transaction.GasLimit = &gasLimit
	transaction.GasPrice = gasPrice
	transaction.Status = TransactionStatusSending
}

func (transaction *Transaction) Sent(trx *types.Transaction) {
	transaction.Status = TransactionStatusSent
	transaction.Hash = trx.Hash().Bytes()
}

//type TransactionUpdate_StatusSent struct {
//	Status    TransactionStatus `json:"status" db:"status"`
//	UpdatedAt *time.Time        `json:"updated_at,omitempty" db:"updated_at"`
//	Hash      []byte            `json:"hash" db:"hash"`
//}
//
//type TransactionUpdate_StatusConfirmed struct {
//	Status    TransactionStatus `json:"status" db:"status"`
//	UpdatedAt *time.Time        `json:"updated_at,omitempty" db:"updated_at"`
//}
//
//type TransactionUpdate_StatusFailed struct {
//	Status    TransactionStatus `json:"status" db:"status"`
//	UpdatedAt *time.Time        `json:"updated_at,omitempty" db:"updated_at"`
//	Error     *string           `json:"error" db:"error"`
//}
//
//func NewTransactionID() ID {
//	return NewID("trx")
//}
//
//func (transaction *Transaction) Table() string {
//	return TransactionsTable
//}
//
//func (transaction *Transaction) Create(db *goqu.Database) error {
//	_, err := db.Insert(transaction.Table()).
//		Returning(goqu.C("id")).
//		Rows(transaction).Executor().Exec()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (transaction *Transaction) Update(
//	db *goqu.Database,
//	update interface{},
//) error {
//	facts := karma.
//		Describe("transaction", transaction).
//		Describe("update", update)
//
//	var (
//		query = db.Update(transaction.Table()).Set(update)
//		where = goqu.Ex{
//			"id": transaction.ID,
//		}
//	)
//
//	switch update.(type) {
//	case TransactionUpdate_StatusSending:
//		where["status"] = TransactionStatusPending
//	case TransactionUpdate_StatusSent:
//		where["status"] = TransactionStatusSending
//	case TransactionUpdate_StatusFailed:
//		// pass
//	case TransactionUpdate_StatusConfirmed:
//		where["status"] = TransactionStatusSent
//	default:
//		return facts.Reason(
//			"unsupported update transition",
//		)
//	}
//
//	result, err := query.Where(where).Executor().Exec()
//	if err != nil {
//		return facts.
//			Format(
//				err,
//				"unable to update transaction",
//			)
//	}
//
//	affected, err := result.RowsAffected()
//	if err != nil {
//		return facts.Format(
//			err,
//			"unable to retrieve affected rows count",
//		)
//	}
//
//	if affected != 1 {
//		return facts.Reason(
//			fmt.Errorf(
//				"update expected to affect exactly 1 row, but %d affected",
//				affected,
//			),
//		)
//	}
//
//	return nil
//}
