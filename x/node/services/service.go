package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/dun-io/imhub/x/node/keeper"
	"github.com/dun-io/imhub/x/node/services/v1"
	v1types "github.com/dun-io/imhub/x/node/types/v1"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1.NewMsgServerImpl(k))

	v1types.RegisterQueryServer(configurator.QueryServer(), v1.NewQueryServerImpl(k))
}
