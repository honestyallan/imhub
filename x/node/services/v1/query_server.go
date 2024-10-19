package v1

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dun-io/imhub/x/node/keeper"
	"github.com/dun-io/imhub/x/node/types/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ v1.QueryServer = (*queryServer)(nil)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v1.QueryServer {
	return &queryServer{k}
}

func (q *queryServer) QueryNodes(c context.Context, req *v1.QueryNodesRequest) (res *v1.QueryNodesResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryNodes(ctx, req)
}
