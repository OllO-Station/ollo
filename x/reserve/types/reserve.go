package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (metadata DenomWhitelist) Validate() error {
	for _, addr := range metadata.Addresses {
		if addr != "" {
			_, err := sdk.AccAddressFromBech32(addr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
