package swap

import (
	"context"
	"fmt"
	"sync"

	"github.com/zuni-lab/yexus-api/pkg/db"
)

type PoolExtendedInfo struct {
	db.PoolDetailsRow
	USDPrice string
}

type poolsInfo struct {
	pools map[string]*PoolExtendedInfo
	mu    sync.RWMutex
}

var PoolInfo *poolsInfo

func InitPoolInfo() {
	PoolInfo = &poolsInfo{
		pools: make(map[string]*PoolExtendedInfo),
		mu:    sync.RWMutex{},
	}
}

func (p *poolsInfo) getTokenInfo(ctx context.Context, poolAddress string) (*PoolExtendedInfo, error) {

	// Check cache first
	p.mu.RLock()
	info, exists := p.pools[poolAddress]
	p.mu.RUnlock()

	if exists {
		return info, nil
	}

	// Get pool info from database
	pool, err := db.DB.PoolDetails(ctx, poolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool: %w", err)
	}

	usdPrice := ""

	extendedPool := PoolExtendedInfo{
		PoolDetailsRow: pool,
		USDPrice:       usdPrice,
	}

	p.mu.Lock()
	p.pools[poolAddress] = &extendedPool
	p.mu.Unlock()

	return &extendedPool, nil
}

func (p *poolsInfo) updateUsdPrice(poolAddress string, usdPrice string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.pools[poolAddress].USDPrice = usdPrice
	return nil
}

type PricesInfo struct {
	TokenName string `json:"token_name"`
	Price     string `json:"price"`
}

func (p *poolsInfo) GetPrices(ctx context.Context) ([]*PricesInfo, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	prices := make([]*PricesInfo, 0)
	for _, pool := range p.pools {
		name := pool.Token0Symbol
		if pool.Token0IsStable {
			name = pool.Token1Symbol
		}

		prices = append(prices, &PricesInfo{
			TokenName: name,
			Price:     pool.USDPrice,
		})
	}
	return prices, nil
}
