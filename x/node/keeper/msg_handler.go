package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	base "github.com/dun-io/imhub/types"
	v1base "github.com/dun-io/imhub/types/v1"
	"github.com/dun-io/imhub/x/node/types"
	"github.com/dun-io/imhub/x/node/types/v1"
	"time"
)

func (k *Keeper) HandleMsgRegisterNode(ctx sdk.Context, msg *v1.MsgRegisterNodeRequest) (*v1.MsgRegisterNodeResponse, error) {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr := base.NodeAddress(accAddr.Bytes())
	if k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorDuplicateNode(nodeAddr)
	}

	deposit := k.Deposit(ctx)
	if err := k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	node := v1.Node{
		Name:        msg.Name,
		Address:     nodeAddr.String(),
		PubKey:      msg.PubKey,
		StakeAmount: 0,
		InactiveAt:  time.Time{},
		Status:      v1base.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	k.SetNode(ctx, node)
	k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)

	return &v1.MsgRegisterNodeResponse{}, nil
}
