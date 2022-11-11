package keeper

import (
	"context"
	"ollo/x/ons/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// "github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/types/query"

)

// Querier is used as Keeper will have duplicate methods if used directly, and gfunc(k Querier) names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

	func(k Querier) OlloName(c context.Context, _ *types.QueryOlloNameRequest) (*types.QueryOlloNameResponse, error) {
    return &types.QueryOlloNameResponse{}, nil
}
	func(k Querier) IbcNames(c context.Context, _ *types.QueryIbcNamesRequest) (*types.QueryIbcNamesResponse, error) {
    return &types.QueryIbcNamesResponse{}, nil
}
	func(k Querier) Threads(c context.Context, _ *types.QueryThreadsRequest) (*types.QueryThreadsResponse, error) {
    return &types.QueryThreadsResponse{}, nil
}
	func(k Querier) Thread(c context.Context, _ *types.QueryThreadRequest) (*types.QueryThreadResponse, error) {
    return &types.QueryThreadResponse{}, nil
}
	func(k Querier) AddressNames(c context.Context, _ *types.QueryGetAddressNamesRequest) (*types.QueryGetAddressNamesResponse, error) {
    return &types.QueryGetAddressNamesResponse{}, nil
}
	func(k Querier) NameInfo(c context.Context, _ *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
}
	func(k Querier) NameThreads(c context.Context, _ *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
}
	func(k Querier) NameThread(c context.Context, _ *types.QueryGetNameThreadRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
	}
	func(k Querier) NameThreadReply(c context.Context, _ *types.QueryGetNameThreadReplyRequest) (*types.QueryGetNameThreadReplyResponse, error) {
    return &types.QueryGetNameThreadReplyResponse{}, nil
	}
	func(k Querier) NameTags(c context.Context, _ *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
	}
	func(k Querier) NameOffers(c context.Context, _ *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
	}
	func(k Querier) NameLoans(c context.Context, _ *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
	}
	func(k Querier) GroupNames(c context.Context, _ *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
	}
	func(k Querier) GroupName(c context.Context, _ *types.QueryGetGroupRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
	}
	func(k Querier) GroupTags(c context.Context, _ *types.QueryGetGroupTagsRequest) (*types.QueryGetGroupTagsResponse, error) {
    return &types.QueryGetGroupTagsResponse{}, nil
	}
	func(k Querier) Tags(c context.Context, _ *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
	}

	func(k Querier) Tag(c context.Context, req *types.QueryGetTagRequest) (*types.QueryGetNameResponse, error) {
    return &types.QueryGetNameResponse{Name: types.Name{}}, nil
	}

  // Queries a Name by index.

// // Airdrops queries all the existing airdrops.
// func (k Querier) Names(c context.Context, req *types.QueryAllNames) (*types.QueryAllNames, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "empty request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(c)
//   return k.GetNames(ctx)
// }
func (k Querier) Names(c context.Context, req *types.QueryGetNamesRequest) (*types.QueryGetNamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// ctx := sdk.UnwrapSDKContext(c)
  return &types.QueryGetNamesResponse{Names: []types.Name{}}, nil
}
func (k Querier) NameId(c context.Context, req *types.QueryGetNameIdRequest) (*types.QueryGetNameIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
  res, err := k.GetName(ctx, &types.QueryGetNameRequest{Name: ""})
  return &types.QueryGetNameIdResponse{res.Name}, err
}
// 	airdropStore := prefix.NewStore(store, types.AirdropKeyPrefix)

// 	airdrops := []types.Airdrop{}
// 	pageRes, err := query.FilteredPaginate(airdropStore, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
// 		var airdrop types.Airdrop
// 		k.cdc.MustUnmarshal(value, &airdrop)

// 		if accumulate {
// 			airdrops = append(airdrops, airdrop)
// 		}

// 		return true, nil
// 	})
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, err.Error())
// 	}

// 	return &types.QueryAirdropsResponse{Airdrops: airdrops, Pagination: pageRes}, nil
// }

// // Airdrop queries the specific airdrop.
// func (k Querier) Airdrop(c context.Context, req *types.QueryAirdropRequest) (*types.QueryAirdropResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "empty request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(c)

// 	airdrop, found := k.Keeper.GetAirdrop(ctx, req.AirdropId)
// 	if !found {
// 		return nil, status.Error(codes.NotFound, "airdrop not found")
// 	}

// 	return &types.QueryAirdropResponse{Airdrop: airdrop}, nil
// }

// // ClaimRecord queries the specific claim record.
// func (k Querier) ClaimRecord(c context.Context, req *types.QueryClaimRecordRequest) (*types.QueryClaimRecordResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(c)

// 	recipientAddr, err := sdk.AccAddressFromBech32(req.Recipient)
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid address")
// 	}

// 	record, found := k.GetClaimRecordByRecipient(ctx, req.AirdropId, recipientAddr)
// 	if !found {
// 		return nil, status.Error(codes.NotFound, "claim record not found")
// 	}

// 	return &types.QueryClaimRecordResponse{ClaimRecord: record}, nil
// }
