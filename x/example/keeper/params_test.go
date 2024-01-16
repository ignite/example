package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/example/testutil/keeper"
	"github.com/ignite/example/x/example/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ExampleKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
