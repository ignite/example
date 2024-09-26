package example_test

import (
	"testing"

	keepertest "github.com/ignite/example/testutil/keeper"
	"github.com/ignite/example/testutil/nullify"
	example "github.com/ignite/example/x/example/module"
	"github.com/ignite/example/x/example/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExampleKeeper(t)
	example.InitGenesis(ctx, k, genesisState)
	got := example.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
