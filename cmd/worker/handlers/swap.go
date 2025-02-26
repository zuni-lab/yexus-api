package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/evm"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

type swapHandler struct {
	tokens map[string]*db.PoolDetailsRow // Cache token info by pool address
}

var _ evm.SwapHandler = &swapHandler{}

func NewSwapHandler() *swapHandler {
	return &swapHandler{
		tokens: make(map[string]*db.PoolDetailsRow),
	}
}

func (h *swapHandler) HandleSwap(ctx context.Context, event *evm.UniswapV3Swap) error {
	log.Info().
		Str("pool", event.Raw.Address.Hex()).
		Msg("Handling swap event")

	poolAddress := event.Raw.Address.Hex()

	// Get or load token info
	tokenInfo, err := h.getTokenInfo(ctx, poolAddress)
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
		Msg("Swap event")

	// Calculate price
	price := utils.CalculatePrice(
		event.SqrtPriceX96,
		uint8(tokenInfo.Token0Decimals),
		uint8(tokenInfo.Token1Decimals),
		tokenInfo.Token0IsStable,
	)

	if price == nil {
		return fmt.Errorf("failed to calculate price for pool %s", poolAddress)
	}

	// Convert big.Float to pgtype.Numeric
	var priceNumeric pgtype.Numeric
	priceStr := price.Text('f', 18) // Use 18 decimal places for precision

	if err := priceNumeric.Scan(priceStr); err != nil {
		return fmt.Errorf("failed to convert price to numeric: %w", err)
	}

	log.Info().
		Str("pool", poolAddress).
		Any("price", priceNumeric).
		Msg("Calculated price")

	// Store price in database
	_, err = db.DB.CreatePrice(ctx, db.CreatePriceParams{
		PoolID:   poolAddress,
		PriceUsd: priceNumeric,
	})

	if err != nil {
		return fmt.Errorf("failed to store price: %w", err)
	}

	log.Info().
		Str("pool", poolAddress).
		Str("price", priceStr).
		Msg("Successfully processed swap event and stored price")

	return nil
}

func (h *swapHandler) getTokenInfo(ctx context.Context, poolAddress string) (*db.PoolDetailsRow, error) {
	// Check cache first
	if info, exists := h.tokens[poolAddress]; exists {
		return info, nil
	}

	poolAddress = strings.ToLower(poolAddress)

	// Get pool info from database
	pool, err := db.DB.PoolDetails(ctx, poolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool: %w", err)
	}

	h.tokens[poolAddress] = &pool

	return h.tokens[poolAddress], nil
}
