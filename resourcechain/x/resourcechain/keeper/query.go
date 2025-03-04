package keeper

import (
	"resourcechain/x/resourcechain/types"
)

var _ types.QueryServer = Keeper{}
