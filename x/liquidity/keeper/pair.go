package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"ollo/x/liquidity/types"
)

// getNextPairIdWithUpdate increments pair id by one and set it.
func (k Keeper) getNextPairIdWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastPairId(ctx) + 1
	k.SetLastPairId(ctx, id)
	return id
}

// getNextOrderIdWithUpdate increments the pair's last order id and returns it.
func (k Keeper) getNextOrderIdWithUpdate(ctx sdk.Context, pair types.Pair) uint64 {
	id := pair.LastOrderId + 1
	pair.LastOrderId = id
	k.SetPair(ctx, pair)
	return id
}

// ValidateMsgCreatePair validates types.MsgCreatePair.
func (k Keeper) ValidateMsgCreatePair(ctx sdk.Context, msg *types.MsgCreatePair) error {
	if _, found := k.GetPairByDenoms(ctx, msg.BaseDenom, msg.QuoteDenom); found {
		return types.ErrPairAlreadyExists
	}
	return nil
}

// CreatePair handles types.MsgCreatePair and creates a pair.
func (k Keeper) CreatePair(ctx sdk.Context, msg *types.MsgCreatePair) (types.Pair, error) {
	if err := k.ValidateMsgCreatePair(ctx, msg); err != nil {
		return types.Pair{}, err
	}

	feeCollector := k.GetFeeCollector(ctx)
	pairCreationFee := k.GetPairCreationFee(ctx)

	// Send the pair creation fee to the fee collector.
	addr, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		return types.Pair{}, err
	}
	if err := k.bankKeeper.SendCoins(ctx, addr, feeCollector, pairCreationFee); err != nil {
		return types.Pair{}, sdkerrors.Wrap(err, "insufficient pair creation fee")
	}

	id := k.getNextPairIdWithUpdate(ctx)
	pair := types.NewPair(id, msg.BaseDenom, msg.QuoteDenom)
	k.SetPair(ctx, pair)
	k.SetPairIndex(ctx, pair.BaseDenom, pair.QuoteDenom, pair.Id)
	k.SetPairLookupIndex(ctx, pair.BaseDenom, pair.QuoteDenom, pair.Id)
	k.SetPairLookupIndex(ctx, pair.QuoteDenom, pair.BaseDenom, pair.Id)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatePair,
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyBaseDenom, msg.BaseDenom),
			sdk.NewAttribute(types.AttributeKeyQuoteDenom, msg.QuoteDenom),
			sdk.NewAttribute(types.AttributeKeyPairId, strconv.FormatUint(pair.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyEscrowAddress, pair.EscrowAddress),
		),
	})

	return pair, nil
}
