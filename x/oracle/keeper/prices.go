package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"
	"ollo/x/oracle/types"
)

// SetPricesResult saves the Prices result
func (k Keeper) SetPricesResult(ctx sdk.Context, requestID types.OracleRequestID, result types.PricesResult) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PricesResultStoreKey(requestID), k.cdc.MustMarshal(&result))
}

// GetPricesResult returns the Prices by requestId
func (k Keeper) GetPricesResult(ctx sdk.Context, id types.OracleRequestID) (types.PricesResult, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.PricesResultStoreKey(id))
	if bz == nil {
		return types.PricesResult{}, sdkerrors.Wrapf(types.ErrSample,
			"GetResult: Result for request ID %d is not available.", id,
		)
	}
	var result types.PricesResult
	k.cdc.MustUnmarshal(bz, &result)
	return result, nil
}

// GetLastPricesID return the id from the last Prices request
func (k Keeper) GetLastPricesID(ctx sdk.Context) int64 {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastPricesIDKey))
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
	return intV.GetValue()
}

// SetLastPricesID saves the id from the last Prices request
func (k Keeper) SetLastPricesID(ctx sdk.Context, id types.OracleRequestID) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.LastPricesIDKey),
		k.cdc.MustMarshalLengthPrefixed(&gogotypes.Int64Value{Value: int64(id)}))
}
