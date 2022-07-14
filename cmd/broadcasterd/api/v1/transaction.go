package v1

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"

	"broadcaster/internal/logging"
	"broadcaster/internal/model"
	"broadcaster/internal/repository/transactions"
	"broadcaster/internal/response"
)

func TransactionInfo(w http.ResponseWriter, r *http.Request) {
	hashSource := common.HexToHash(mux.Vars(r)["hashSource"])
	chainIdFrom, _ := strconv.ParseInt(mux.Vars(r)["chainIdFrom"], 10, 32)

	transaction := &model.Transaction{}

	if _, err := transactions.Query(r.Context()).
		Where("hash_source = ? AND chain_id_from = ?", hashSource, chainIdFrom).
		Exec(r.Context(), transaction); errors.Is(err, sql.ErrNoRows) || transaction.Hash == nil {
		response.NotFoundError().Write(w, r)
		return
	} else if err != nil {
		panic(err)
	}
	logging.CtxGet(r.Context()).Info(hashSource.String())

	response.JSON(w, map[string]interface{}{
		"hash":    common.BytesToHash(transaction.Hash).String(),
		"chainId": transaction.ChainID,
	})
}
