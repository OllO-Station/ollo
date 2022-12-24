package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"ollo/x/grants/simulation"
)

func TestParamChanges(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"fundraising/AuctionCreationFee", "AuctionCreationFee", "[{\"denom\":\"stake\",\"amount\":\"98498081\"}]", "fundraising"},
		{"fundraising/ExtendedPeriod", "ExtendedPeriod", "7", "fundraising"},
	}

	paramChanges := simulation.ParamChanges(r)
	require.Len(t, paramChanges, 2)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}
