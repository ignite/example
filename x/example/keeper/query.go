package keeper

import (
	"github.com/ignite/example/x/example/types"
)

var _ types.QueryServer = Keeper{}
