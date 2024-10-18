package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dun-io/imhub/x/node/types"
	"github.com/dun-io/imhub/x/node/types/v1"

	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) (params v1.Params) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params v1.Params) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)

	return nil
}
func (k *Keeper) IsValidGigabytePrices(ctx sdk.Context, prices sdk.Coins) bool {
	minPrices := k.MinGigabytePrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

// IsValidHourlyPrices checks if the provided hourly prices are valid based on the minimum prices defined in the module's parameters.
func (k *Keeper) IsValidHourlyPrices(ctx sdk.Context, prices sdk.Coins) bool {
	minPrices := k.MinHourlyPrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}
func (k *Keeper) MinGigabytePrices(ctx sdk.Context) sdk.Coins {
	return k.GetParams(ctx).MinGigabytePrices
}
func (k *Keeper) MinHourlyPrices(ctx sdk.Context) sdk.Coins {
	return k.GetParams(ctx).MinHourlyPrices
}
func (k *Keeper) Deposit(ctx sdk.Context) sdk.Coin {
	return k.GetParams(ctx).Deposit
}
