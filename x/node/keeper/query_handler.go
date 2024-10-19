package keeper

import (
	"cosmossdk.io/store/prefix"
	v1base "github.com/dun-io/imhub/types/v1"
	"github.com/dun-io/imhub/x/node/types"
	"github.com/dun-io/imhub/x/node/types/v1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HandleQueryNodes(ctx sdk.Context, req *v1.QueryNodesRequest) (*v1.QueryNodesResponse, error) {
	var (
		items     []v1.Node
		keyPrefix []byte
	)

	switch req.Status {
	case v1base.StatusActive:
		keyPrefix = types.ActiveNodeKeyPrefix
	case v1base.StatusInactive:
		keyPrefix = types.InactiveNodeKeyPrefix
	default:
		keyPrefix = types.NodeKeyPrefix
	}

	store := prefix.NewStore(k.Store(ctx), keyPrefix)
	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v1.Node
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryNodesResponse{Nodes: items, Pagination: pagination}, nil
}
