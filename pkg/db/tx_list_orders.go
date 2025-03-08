package db

import (
	"context"
	"github.com/jinzhu/copier"
)

type ListOrdersByWalletResult struct {
	Orders []GetOrdersByWalletRow
	Count  int64
}

func (store *SqlStore) ListOrdersByWalletTx(ctx context.Context, params GetOrdersByWalletParams) (*ListOrdersByWalletResult, error) {
	var result *ListOrdersByWalletResult

	err := store.execTx(ctx, func(q *Queries) error {
		var (
			err         error
			countParams CountOrdersByWalletParams
		)

		orders, err := DB.GetOrdersByWallet(ctx, params)
		if err != nil {
			return err
		}

		err = copier.Copy(&countParams, &params)
		if err != nil {
			return err
		}

		count, err := DB.CountOrdersByWallet(ctx, countParams)
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
