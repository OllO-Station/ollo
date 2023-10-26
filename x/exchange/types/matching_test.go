package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/testutil"
	utils "github.com/ollo-station/ollo/x/ollo/types"
	"github.com/ollo-station/ollo/x/exchange/types"
)

func TestPayReceiveDenoms(t *testing.T) {
	payDenom, receiveDenom := types.PayReceiveDenoms("uollo", "uusd", true)
	require.Equal(t, "uusd", payDenom)
	require.Equal(t, "uollo", receiveDenom)
	payDenom, receiveDenom = types.PayReceiveDenoms("uollo", "uusd", false)
	require.Equal(t, "uollo", payDenom)
	require.Equal(t, "uusd", receiveDenom)
}

func TestFillMemOrderBasic(t *testing.T) {
	market := types.NewMarket(
		1, "uollo", "uusd", utils.ParseDec("-0.0015"), utils.ParseDec("0.003"), utils.ParseDec("0.5"))
	ctx := types.NewMatchingContext(market, false)

	order := newUserMemOrder(1, true, utils.ParseDec("1.3"), sdk.NewDec(10_000000), sdk.NewDec(9_000000))
	ctx.FillOrder(order, sdk.NewDec(5_000000), utils.ParseDec("1.25"), true)

	require.True(t, order.IsMatched())
	testutil.AssertEqual(t, sdk.NewDec(6_240625), order.Paid())
	testutil.AssertEqual(t, sdk.NewDec(5_000000), order.Received())
	testutil.AssertEqual(t, sdk.NewDec(-9375), order.Fee())

	order = newUserMemOrder(2, false, utils.ParseDec("1.2"), sdk.NewDec(10_000000), sdk.NewDec(9_000000))
	ctx.FillOrder(order, sdk.NewDec(5_000000), utils.ParseDec("1.25"), false)

	testutil.AssertEqual(t, sdk.NewDec(5_000000), order.Paid())
	testutil.AssertEqual(t, sdk.NewDec(6_231250), order.Received())
	testutil.AssertEqual(t, sdk.NewDec(18750), order.Fee())
}
