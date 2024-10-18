package v1

import (
	"github.com/dun-io/imhub/x/node/keeper"
	"github.com/dun-io/imhub/x/node/types/v1"
)

var _ v1.QueryServer = keeper.Keeper{}
