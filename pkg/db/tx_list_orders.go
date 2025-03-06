package db

import (
	"context"
)

type ListOrdersByWalletResult struct {
	Orders []Order
	Count  int64
}

func (store *SqlStore) ListOrdersByWalletTx(ctx context.Context, params GetOrdersByWalletParams) (*ListOrdersByWalletResult, error) {
	var result *ListOrdersByWalletResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		orders, err := DB.GetOrdersByWallet(ctx, params)
		if err != nil {
			return err
		}

		count, err := DB.CountOrdersByWallet(ctx, CountOrdersByWalletParams{
			Wallet: params.Wallet,
			Types:  params.Types,
			Status: params.Status,
			Side:   params.Side,
		})
		if err != nil {
			return err
		}

		result = &ListOrdersByWalletResult{
			Orders: orders,
			Count:  count,
		}

		return nil
	})

	return result, err
}
