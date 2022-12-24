package types_test

import (
	"testing"

	"ollo/x/grants/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
)

func TestValidate_AllowedBidder(t *testing.T) {
	testBidderAddr := sdk.AccAddress(crypto.AddressHash([]byte("TestBidder")))

	testCases := []struct {
		allowedBidder types.AllowedBidder
		expectedErr   bool
	}{
		{
			types.NewAllowedBidder(testBidderAddr, sdk.NewInt(100_000_000)),
			false,
		},
		{
			types.NewAllowedBidder(sdk.AccAddress{}, sdk.NewInt(100_000_000)),
			true,
		},
		{
			types.NewAllowedBidder(testBidderAddr, sdk.NewInt(0)),
			true,
		},
		{
			types.NewAllowedBidder(testBidderAddr, sdk.ZeroInt()),
			true,
		},
	}

	for _, tc := range testCases {
		err := tc.allowedBidder.Validate()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.Equal(t, testBidderAddr, tc.allowedBidder.GetBidder())
			require.NoError(t, err)
		}
	}
}
