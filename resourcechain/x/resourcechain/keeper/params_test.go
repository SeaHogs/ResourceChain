package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "resourcechain/testutil/keeper"
	"resourcechain/x/resourcechain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ResourcechainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
