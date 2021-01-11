package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreateAdd{}, "bling/CreateAdd", nil)
cdc.RegisterConcrete(&MsgUpdateAdd{}, "bling/UpdateAdd", nil)
cdc.RegisterConcrete(&MsgDeleteAdd{}, "bling/DeleteAdd", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateAdd{},
	&MsgUpdateAdd{},
	&MsgDeleteAdd{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
