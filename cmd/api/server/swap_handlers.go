package server

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/evm"
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

	return nil
}
