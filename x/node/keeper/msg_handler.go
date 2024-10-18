package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dun-io/imhub/x/node/types/v1"
)

func (k *Keeper) HandleMsgRegisterNode(ctx sdk.Context, msg *v1.MsgRegisterNodeRequest) (*v1.MsgRegisterNodeResponse, error) {

	return &v1.MsgRegisterNodeResponse{}, nil
}
