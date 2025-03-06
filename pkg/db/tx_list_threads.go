package db

import (
	"context"
)

type ListThreadsByWalletResult struct {
	Threads []ChatThread
	Count   int64
}

func (store *SqlStore) ListThreadsTx(ctx context.Context, params GetChatThreadsParams) (*ListThreadsByWalletResult, error) {
	var result *ListThreadsByWalletResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		threads, err := DB.GetChatThreads(ctx, params)
		if err != nil {
			return err
		}

		count, err := DB.CountChatThreads(ctx, params.UserAddress)
		if err != nil {
			return err
		}

		result = &ListThreadsByWalletResult{
			Threads: threads,
			Count:   count,
		}
		return nil
	})

	return result, err
}
