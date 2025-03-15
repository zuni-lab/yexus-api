package swap

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/yexus-api/internal/orders/services"
	"github.com/zuni-lab/yexus-api/pkg/evm"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

type swapHandler struct {
}

var _ evm.SwapHandler = &swapHandler{}

func NewSwapHandler() *swapHandler {
	return &swapHandler{}
}

func (h *swapHandler) HandleSwap(ctx context.Context, event *evm.UniswapV3Swap) error {
	defer utils.Recover("SwapHandler", event, "Failed to handle Swap event")

	poolAddress := utils.NormalizeAddress(event.Raw.Address.Hex())

	// Get or load token info
	tokenInfo, err := PoolInfo.getTokenInfo(ctx, poolAddress)
	if err != nil {
		return fmt.Errorf("failed to get token info: %w", err)
	}

	// Skip if neither token is USD-based
	if !tokenInfo.Token0IsStable && !tokenInfo.Token1IsStable {
		log.Debug().
			Str("pool", poolAddress).
			Msg("Skipping price calculation for non-USD pair")
		return nil
	}

	log.Info().
		Str("pool", poolAddress).
		Str("sqrtPriceX96", event.SqrtPriceX96.String()).
		Msg("Handling Swap event")

	// Calculate price
	price := utils.CalculatePrice(
		event.SqrtPriceX96,
		tokenInfo.Token0Decimals,
		tokenInfo.Token1Decimals,
		tokenInfo.Token0IsStable,
	)

	if price == nil {
		return fmt.Errorf("failed to calculate price for pool %s", poolAddress)
	}

	err = PoolInfo.updateUsdPrice(poolAddress, price.Text('f', -1))
	if err != nil {
		return fmt.Errorf("failed to update usd price for pool %s", poolAddress)
	}

	_, err = services.MatchOrder(ctx, price)
	if err != nil {
		log.Warn().Any("pool", poolAddress).Err(err).Msgf("⚠️ [SwapHandler] Failed to match order for pool %s, at price %s", event.Raw.Address.Hex(), price.String())
		return nil
	}

	log.Info().Any("pool", poolAddress).Msgf("✅ [SwapHandler] Successfully matched order for pool %s, at price %s", event.Raw.Address.Hex(), price.String())
	return nil
}
