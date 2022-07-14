package transactions

import (
	"broadcaster/internal/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/uptrace/bun"
)

func Status(status ...model.TransactionStatus) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where(`"transaction"."status" IN (?)`, bun.In(status))
	}
}

func ID(id model.ID) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where(`"transaction"."id" = ?`, id)
	}
}

func ByRequest(chainId int64, bridgeAddress common.Address, receiveSide common.Address, callData []byte) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where(`"transaction"."chain_id" = ? AND "transaction"."bridge_address" = ? AND "transaction"."receive_side" = ? AND "transaction"."call_data" = ?`,
			chainId,
			bridgeAddress,
			receiveSide,
			callData,
		)
	}
}
