package v1

import (
	"broadcaster/internal/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	handlers.Default(router)
	router.HandleFunc("/", Status).Methods("GET")
	router.HandleFunc("/v1/status", Status).Methods("GET")
	router.HandleFunc("/v1/broadcast/from/{chainIdFrom:[0-9]+}/to/{chainId:[0-9]+}/{bridge:[a-zA-Z0-9]+}", Broadcast).Methods("POST")
	router.HandleFunc("/v1/transactions/{chainIdFrom:[0-9]+}/{hashSource:0x[a-zA-Z0-9]+}", TransactionInfo).Methods("GET")
	return router
}
