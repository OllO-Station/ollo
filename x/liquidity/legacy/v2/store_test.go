package v2_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	chain "github.com/ollo-station/ollo/app"
	utils "github.com/ollo-station/ollo/types"

	v1liquidity "github.com/ollo-station/ollo/x/liquidity/legacy/v1"
	v2liquidity "github.com/ollo-station/ollo/x/liquidity/legacy/v2"
	"github.com/ollo-station/ollo/x/liquidity/types"
)

func TestMigratePool(t *testing.T) {
	cdc := chain.MakeEncodingConfig().Marshaler
	storeKey := sdk.NewKVStoreKey("liquidity")
	ctx := testutil.DefaultContext(storeKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(storeKey)

	oldPool := v1liquidity.Pool{
		Id:                    1,
		PairId:                1,
		ReserveAddress:        utils.TestAddress(0).String(),
		PoolCoinDenom:         "pool1",
		LastDepositRequestId:  2,
		LastWithdrawRequestId: 3,
		Disabled:              true,
	}
	oldPoolValue := cdc.MustMarshal(&oldPool)
	key := types.GetPoolKey(oldPool.Id)
	store.Set(key, oldPoolValue)

	require.NoError(t, v2liquidity.MigrateStore(ctx, storeKey, cdc))

	newPool := types.Pool{
		Type:                  types.PoolTypeBasic,
		Id:                    1,
		PairId:                1,
		Creator:               "",
		ReserveAddress:        utils.TestAddress(0).String(),
		PoolCoinDenom:         "pool1",
		MinPrice:              nil,
		MaxPrice:              nil,
		LastDepositRequestId:  2,
		LastWithdrawRequestId: 3,
		Disabled:              true,
	}
	newPoolValue := cdc.MustMarshal(&newPool)
	require.Equal(t, newPoolValue, store.Get(key))
}
