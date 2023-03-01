package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"

	"github.com/ollo-station/ollo/x/grants/types"
)

func TestConvertToSellingAmount(t *testing.T) {
	payingCoinDenom := "denom2" // auction paying coin denom

	testCases := []struct {
		bid         types.Bid
		expectedAmt sdk.Int
	}{
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.5"),
				Coin:  sdk.NewCoin("denom1", sdk.NewInt(100_000)),
			},
			sdk.NewInt(100_000),
		},
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.5"),
				Coin:  sdk.NewCoin("denom2", sdk.NewInt(100_000)),
			},
			sdk.NewInt(200_000),
		},
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.1"),
				Coin:  sdk.NewCoin("denom1", sdk.NewInt(100_000)),
			},
			sdk.NewInt(100_000),
		},
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.1"),
				Coin:  sdk.NewCoin("denom2", sdk.NewInt(100_000)),
			},
			sdk.NewInt(1_000_000),
		},
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("3"),
				Coin:  sdk.NewCoin("denom2", sdk.NewInt(4)),
			},
			sdk.NewInt(1),
		},
	}

	for _, tc := range testCases {
		sellingAmt := tc.bid.ConvertToSellingAmount(payingCoinDenom)
		require.Equal(t, tc.expectedAmt, sellingAmt)
	}
}

func TestConvertToPayingAmount(t *testing.T) {
	payingCoinDenom := "denom2" // auction paying coin denom

	testCases := []struct {
		bid         types.Bid
		expectedAmt sdk.Int
	}{
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.5"),
				Coin:  sdk.NewCoin("denom1", sdk.NewInt(100_000)),
			},
			sdk.NewInt(50_000),
		},
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.5"),
				Coin:  sdk.NewCoin("denom2", sdk.NewInt(100_000)),
			},
			sdk.NewInt(100_000),
		},
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.1"),
				Coin:  sdk.NewCoin("denom1", sdk.NewInt(100_000)),
			},
			sdk.NewInt(10_000),
		},
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.1"),
				Coin:  sdk.NewCoin("denom2", sdk.NewInt(100_000)),
			},
			sdk.NewInt(100_000),
		},
		{
			types.Bid{
				Price: sdk.MustNewDecFromStr("0.33"),
				Coin:  sdk.NewCoin("denom1", sdk.NewInt(100_000)),
			},
			sdk.NewInt(33000),
		},
	}

	for _, tc := range testCases {
		payingAmt := tc.bid.ConvertToPayingAmount(payingCoinDenom)
		require.Equal(t, tc.expectedAmt, payingAmt)
	}
}

func TestSetMatched(t *testing.T) {
	bidder := sdk.AccAddress(crypto.AddressHash([]byte("Bidder")))

	bid := types.NewBid(
		1,
		bidder,
		1,
		types.BidTypeFixedPrice,
		sdk.MustNewDecFromStr("0.5"),
		sdk.NewCoin("denom1", sdk.NewInt(100_000)),
		false,
	)
	require.False(t, bid.IsMatched)
	require.Equal(t, bidder, bid.GetBidder())

	bid.SetMatched(true)
	require.True(t, bid.IsMatched)
}
