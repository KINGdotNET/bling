package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kingdotnet/bling/x/bling/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetAddCount get the total number of add
func (k Keeper) GetAddCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddCountKey))
	byteKey := types.KeyPrefix(types.AddCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetAddCount set the total number of add
func (k Keeper) SetAddCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddCountKey))
	byteKey := types.KeyPrefix(types.AddCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateAdd(ctx sdk.Context, msg types.MsgCreateAdd) {
	// Create the add
    count := k.GetAddCount(ctx)
    var add = types.Add{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Post: msg.Post,
        Title: msg.Title,
        Body: msg.Body,
        Ipfs: msg.Ipfs,
        Parent: msg.Parent,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
    key := types.KeyPrefix(types.AddKey + add.Id)
    value := k.cdc.MustMarshalBinaryBare(&add)
    store.Set(key, value)

    // Update add count
    k.SetAddCount(ctx, count+1)
}

func (k Keeper) UpdateAdd(ctx sdk.Context, add types.Add) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	b := k.cdc.MustMarshalBinaryBare(&add)
	store.Set(types.KeyPrefix(types.AddKey + add.Id), b)
}

func (k Keeper) GetAdd(ctx sdk.Context, key string) types.Add {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	var add types.Add
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AddKey + key)), &add)
	return add
}

func (k Keeper) HasAdd(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	return store.Has(types.KeyPrefix(types.AddKey + id))
}

func (k Keeper) GetAddOwner(ctx sdk.Context, key string) string {
    return k.GetAdd(ctx, key).Creator
}

// DeleteAdd deletes a add
func (k Keeper) DeleteAdd(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	store.Delete(types.KeyPrefix(types.AddKey + key))
}

func (k Keeper) GetAllAdd(ctx sdk.Context) (msgs []types.Add) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AddKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Add
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
