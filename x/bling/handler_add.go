package bling

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/kingdotnet/bling/x/bling/types"
	"github.com/kingdotnet/bling/x/bling/keeper"
)

func handleMsgCreateAdd(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateAdd) (*sdk.Result, error) {
	k.CreateAdd(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateAdd(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateAdd) (*sdk.Result, error) {
	var add = types.Add{
		Creator: msg.Creator,
		Id:      msg.Id,
    	Post: msg.Post,
    	Title: msg.Title,
    	Body: msg.Body,
    	Ipfs: msg.Ipfs,
    	Parent: msg.Parent,
	}

    if msg.Creator != k.GetAddOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdateAdd(ctx, add)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteAdd(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteAdd) (*sdk.Result, error) {
    if !k.HasAdd(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if msg.Creator != k.GetAddOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeleteAdd(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
