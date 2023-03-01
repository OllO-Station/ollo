package cli

import (
	"context"

	"github.com/ollo-station/ollo/x/token/types"

	"github.com/cosmos/cosmos-sdk/client"
)

// queryTokenFees retrieves the fees of issuance and minting for the specified symbol
func queryTokenFees(cliCtx client.Context, symbol string) (types.QueryFeesResponse, error) {
	queryClient := types.NewQueryClient(cliCtx)

	resp, err := queryClient.Fees(context.Background(), &types.QueryFeesRequest{Denom: symbol})
	if err != nil {
		return types.QueryFeesResponse{}, err
	}

	return *resp, err
}
