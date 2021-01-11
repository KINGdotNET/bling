package rest

import (
	"net/http"
	"strconv"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/kingdotnet/bling/x/bling/types"
    "github.com/gorilla/mux"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createAddRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Post string `json:"post"`
	Title string `json:"title"`
	Body string `json:"body"`
	Ipfs string `json:"ipfs"`
	Parent string `json:"parent"`
	
}

func createAddHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createAddRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedPost := req.Post
		
		parsedTitle := req.Title
		
		parsedBody := req.Body
		
		parsedIpfs := req.Ipfs
		
		parsedParent := req.Parent
		

		msg := types.NewMsgCreateAdd(
			req.Creator,
			parsedPost,
			parsedTitle,
			parsedBody,
			parsedIpfs,
			parsedParent,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateAddRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Post string `json:"post"`
	Title string `json:"title"`
	Body string `json:"body"`
	Ipfs string `json:"ipfs"`
	Parent string `json:"parent"`
	
}


func updateAddHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        id := mux.Vars(r)["id"]

		var req updateAddRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedPost := req.Post
		
		parsedTitle := req.Title
		
		parsedBody := req.Body
		
		parsedIpfs := req.Ipfs
		
		parsedParent := req.Parent
		

		msg := types.NewMsgUpdateAdd(
			req.Creator,
            id,
			parsedPost,
			parsedTitle,
			parsedBody,
			parsedIpfs,
			parsedParent,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteAddRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
}

func deleteAddHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        id := mux.Vars(r)["id"]

		var req deleteAddRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeleteAdd(
			req.Creator,
            id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
