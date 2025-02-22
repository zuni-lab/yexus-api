package db

import (
	"context"
)

type CreatePoolTxParams struct {
	ID     string
	Token0 Token
	Token1 Token
}

type CreatePoolTxResult struct {
	Pool Pool
}

func (store *sqlStore) CreatePoolTx(ctx context.Context, arg CreatePoolTxParams) (CreatePoolTxResult, error) {
	var result CreatePoolTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Pool, err = q.CreatePool(ctx, CreatePoolParams{
			ID:       arg.ID,
			Token0ID: arg.Token0.ID,
			Token1ID: arg.Token1.ID,
		})

		if err != nil {
			return err
		}

		_, err = q.CreateToken(ctx, CreateTokenParams{
			ID:       arg.Token0.ID,
			Name:     arg.Token0.Name,
			Symbol:   arg.Token0.Symbol,
			Decimals: arg.Token0.Decimals,
		})

		if err != nil {
			return err
		}

		_, err = q.CreateToken(ctx, CreateTokenParams{
			ID:       arg.Token1.ID,
			Name:     arg.Token1.Name,
			Symbol:   arg.Token1.Symbol,
			Decimals: arg.Token1.Decimals,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

type CreateBatchPoolsTxParams struct {
	Pools []CreatePoolTxParams
}

type CreateBatchPoolsTxResult struct {
	Pools []Pool
}

func (store *sqlStore) CreateBatchPoolsTx(ctx context.Context, arg CreateBatchPoolsTxParams) (CreateBatchPoolsTxResult, error) {
	var result CreateBatchPoolsTxResult
	result.Pools = make([]Pool, 0, len(arg.Pools))

	err := store.execTx(ctx, func(q *Queries) error {
		tokenMap := make(map[string]struct{})

		for _, poolParam := range arg.Pools {
			if _, exists := tokenMap[poolParam.Token0.ID]; !exists {
				_, err := q.CreateToken(ctx, CreateTokenParams{
					ID:       poolParam.Token0.ID,
					Name:     poolParam.Token0.Name,
					Symbol:   poolParam.Token0.Symbol,
					Decimals: poolParam.Token0.Decimals,
				})
				if err != nil {
					return err
				}
				tokenMap[poolParam.Token0.ID] = struct{}{}
			}

			if _, exists := tokenMap[poolParam.Token1.ID]; !exists {
				_, err := q.CreateToken(ctx, CreateTokenParams{
					ID:       poolParam.Token1.ID,
					Name:     poolParam.Token1.Name,
					Symbol:   poolParam.Token1.Symbol,
					Decimals: poolParam.Token1.Decimals,
				})
				if err != nil {
					return err
				}
				tokenMap[poolParam.Token1.ID] = struct{}{}
			}
		}

		for _, poolParam := range arg.Pools {
			pool, err := q.CreatePool(ctx, CreatePoolParams{
				ID:       poolParam.ID,
				Token0ID: poolParam.Token0.ID,
				Token1ID: poolParam.Token1.ID,
			})
			if err != nil {
				return err
			}
			result.Pools = append(result.Pools, pool)
		}

		return nil
	})

	return result, err
}
