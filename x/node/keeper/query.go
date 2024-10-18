package keeper

import (
	"github.com/dun-io/imhub/x/node/types/v1"
)

var _ v1.QueryServer = Keeper{}
