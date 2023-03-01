package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	"github.com/ollo-station/ollo/x/reserve/types"
)

// GetWhitelist returns the authority metadata for a specific denom
func (k Keeper) GetWhitelist(ctx sdk.Context, denom string) (types.DenomWhitelist, error) {
	bz := k.GetDenomPrefixStore(ctx, denom).Get([]byte(types.DenomWhitelistKey))

	metadata := types.DenomWhitelist{}
	err := proto.Unmarshal(bz, &metadata)
	if err != nil {
		return types.DenomWhitelist{}, err
	}
	return metadata, nil
}

// setWhitelist stores authority metadata for a specific denom
func (k Keeper) setWhitelist(ctx sdk.Context, denom string, metadata types.DenomWhitelist) error {
	err := metadata.Validate()
	if err != nil {
		return err
	}

	store := k.GetDenomPrefixStore(ctx, denom)

	bz, err := proto.Marshal(&metadata)
	if err != nil {
		return err
	}

	store.Set([]byte(types.DenomWhitelistKey), bz)
	return nil
}

func (k Keeper) setAddresses(ctx sdk.Context, denom string, admin string) error {
	metadata, err := k.GetWhitelist(ctx, denom)
	if err != nil {
		return err
	}

	metadata.Addresses = append(metadata.Addresses, admin)

	return k.setWhitelist(ctx, denom, metadata)
}
