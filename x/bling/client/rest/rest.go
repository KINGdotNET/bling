package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
    // this line is used by starport scaffolding # 1
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers bling-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 3
    r.HandleFunc("/bling/adds/{id}", getAddHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/bling/adds", listAddHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 4
    r.HandleFunc("/bling/adds", createAddHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/bling/adds/{id}", updateAddHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/bling/adds/{id}", deleteAddHandler(clientCtx)).Methods("POST")

}

