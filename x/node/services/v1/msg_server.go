package v1

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dun-io/imhub/x/node/keeper"
	"github.com/dun-io/imhub/x/node/types/v1"
)

var _ v1.MsgServiceServer = (*msgServer)(nil)

type msgServer struct {
	keeper.Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper keeper.Keeper) v1.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

func (m *msgServer) MsgRegisterNode(c context.Context, req *v1.MsgRegisterNodeRequest) (*v1.MsgRegisterNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return m.HandleMsgRegisterNode(ctx, req)
}
