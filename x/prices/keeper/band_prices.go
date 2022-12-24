package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"
	"ollo/x/prices/types"
)

// SetBandPricesResult saves the BandPrices result
func (k Keeper) SetBandPricesResult(ctx sdk.Context, requestID types.OracleRequestID, result types.BandPricesResult) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.BandPricesResultStoreKey(requestID), k.cdc.MustMarshal(&result))
}

// GetBandPricesResult returns the BandPrices by requestId
func (k Keeper) GetBandPricesResult(ctx sdk.Context, id types.OracleRequestID) (types.BandPricesResult, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.BandPricesResultStoreKey(id))
	if bz == nil {
		return types.BandPricesResult{}, sdkerrors.Wrapf(types.ErrSample,
			"GetResult: Result for request ID %d is not available.", id,
		)
	}
	var result types.BandPricesResult
	k.cdc.MustUnmarshal(bz, &result)
	return result, nil
}

// GetLastBandPricesID return the id from the last BandPrices request
func (k Keeper) GetLastBandPricesID(ctx sdk.Context) int64 {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastBandPricesIDKey))
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
	return intV.GetValue()
}

// SetLastBandPricesID saves the id from the last BandPrices request
func (k Keeper) SetLastBandPricesID(ctx sdk.Context, id types.OracleRequestID) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.LastBandPricesIDKey),
		k.cdc.MustMarshalLengthPrefixed(&gogotypes.Int64Value{Value: int64(id)}))
}
