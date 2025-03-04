package keeper

import (
	"resourcechain/x/resource/types"
)

var _ types.QueryServer = Keeper{}
