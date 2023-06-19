package keeper_test

import (
	"testing"

	helpers "github.com/notional-labs/centauri/v3/app/helpers"
	"github.com/notional-labs/centauri/v3/x/transfermiddleware/types"
	"github.com/stretchr/testify/require"
)

func TestTFMInitGenesis(t *testing.T) {
	app := helpers.SetupCentauriAppWithValSet(t)
	ctx := helpers.NewContextForApp(*app)

	tokenInfos := make([]types.ParachainIBCTokenInfo, 1)
	tokenInfos[0] = types.ParachainIBCTokenInfo{
		IbcDenom:    "ibc-test",
		ChannelId:   "channel-0",
		NativeDenom: "pica",
		AssetId:     "1",
	}

	app.TransferMiddlewareKeeper.InitGenesis(ctx, types.GenesisState{
		TokenInfos: tokenInfos,
	})

	info := app.TransferMiddlewareKeeper.GetParachainIBCTokenInfoByNativeDenom(ctx, "pica")
	require.Equal(t, info, app.TransferMiddlewareKeeper.GetParachainIBCTokenInfoByNativeDenom(ctx, "pica"))
	require.Equal(t, "1", info.AssetId)
	require.Equal(t, "pica", info.NativeDenom)
	require.Equal(t, "ibc-test", info.IbcDenom)
	require.Equal(t, "channel-0", info.ChannelId)
}

func TestTFMExportGenesis(t *testing.T) {
	app := helpers.SetupCentauriAppWithValSet(t)
	ctx := helpers.NewContextForApp(*app)

	// default params
	err := app.TransferMiddlewareKeeper.AddParachainIBCInfo(ctx, "ibc-test", "channel-0", "pica", "1")
	require.NoError(t, err)
	err = app.TransferMiddlewareKeeper.AddParachainIBCInfo(ctx, "ibc-test2", "channel-1", "poke", "2")
	require.NoError(t, err)
	genesis := app.TransferMiddlewareKeeper.ExportGenesis(ctx)

	require.Equal(t, "1", genesis.TokenInfos[0].AssetId)
	require.Equal(t, "pica", genesis.TokenInfos[0].NativeDenom)
	require.Equal(t, "channel-0", genesis.TokenInfos[0].ChannelId)
	require.Equal(t, "ibc-test", genesis.TokenInfos[0].IbcDenom)

	require.Equal(t, "2", genesis.TokenInfos[1].AssetId)
	require.Equal(t, "poke", genesis.TokenInfos[1].NativeDenom)
	require.Equal(t, "channel-1", genesis.TokenInfos[1].ChannelId)
	require.Equal(t, "ibc-test2", genesis.TokenInfos[1].IbcDenom)
}
