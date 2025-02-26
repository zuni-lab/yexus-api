package evm

import "context"

type SwapHandler interface {
	HandleSwap(ctx context.Context, event *UniswapV3Swap) error
}
