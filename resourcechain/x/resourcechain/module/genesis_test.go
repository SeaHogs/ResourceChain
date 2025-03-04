package resourcechain_test

import (
	"testing"

	keepertest "resourcechain/testutil/keeper"
	"resourcechain/testutil/nullify"
	resourcechain "resourcechain/x/resourcechain/module"
	"resourcechain/x/resourcechain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ResourcechainKeeper(t)
	resourcechain.InitGenesis(ctx, k, genesisState)
	got := resourcechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
