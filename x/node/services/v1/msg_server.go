package v1

import (
	"github.com/dun-io/imhub/x/node/keeper"
	"github.com/dun-io/imhub/x/node/types/v1"
)

var _ v1.MsgServer = (*msgServer)(nil)

type msgServer struct {
	keeper.Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper keeper.Keeper) v1.MsgServer {
	return &msgServer{Keeper: keeper}
}
