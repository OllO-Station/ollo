package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ollo-station/ollo/x/claim/types"
)

// SetAirdropSupply set airdropSupply in the store
func (k Keeper) SetAirdropSupply(ctx sdk.Context, airdropSupply sdk.Coin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropSupplyKey))
	b := k.cdc.MustMarshal(&airdropSupply)
	store.Set([]byte{0}, b)
}

// GetAirdropSupply returns airdropSupply
func (k Keeper) GetAirdropSupply(ctx sdk.Context) (val sdk.Coin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropSupplyKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAirdropSupply removes the AirdropSupply from the store
func (k Keeper) RemoveAirdropSupply(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropSupplyKey))
	store.Delete([]byte{0})
}

// InitializeAirdropSupply set the airdrop supply in the store and set the module balance
func (k Keeper) InitializeAirdropSupply(ctx sdk.Context, airdropSupply sdk.Coin) error {
	// get the eventual existing balance of the module for the airdrop supply
	moduleBalance := k.bankKeeper.GetBalance(
		ctx,
		k.accountKeeper.GetModuleAddress(types.ModuleName),
		airdropSupply.Denom,
	)

	// if the module has an existing balance, we burn the entire balance
	if moduleBalance.IsPositive() {
		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(moduleBalance)); err != nil {
			return errors.New("can't burn module balance %s", 1001, err.Error())
		}
	}

	// set the module balance with the airdrop supply
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(airdropSupply)); err != nil {
		return errors.New("can't mint airdrop suply into module balance %s", 1001, err.Error())
	}

	k.SetAirdropSupply(ctx, airdropSupply)
	return nil
}

func (k Keeper) EndAirdrop(ctx sdk.Context) error {
	airdropSupply, found := k.GetAirdropSupply(ctx)
	if !found || !airdropSupply.IsPositive() {
		return nil
	}

	decayInfo := k.DecayInformation(ctx)
	if decayInfo.Enabled && ctx.BlockTime().After(decayInfo.DecayEnd) {
		err := k.distrKeeper.FundCommunityPool(
			ctx,
			sdk.NewCoins(airdropSupply),
			k.accountKeeper.GetModuleAddress(types.ModuleName))
		if err != nil {
			return err
		}

		airdropSupply.Amount = sdk.ZeroInt()
		k.SetAirdropSupply(ctx, airdropSupply)
	}

	// TODO
	// handle other options:
	// https://ollo/issues/53
	return nil
}
