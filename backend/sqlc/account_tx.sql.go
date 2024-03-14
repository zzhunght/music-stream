package sqlc

import (
	"context"

	"github.com/rs/zerolog/log"
)

type CreateAccountWithTxParams struct {
	Params        CreateAccountParams
	AfterFunction func(string) error
}

func (store *SQLStore) CreateAccountWithTx(ctx context.Context, arg CreateAccountWithTxParams) (CreateAccountRow, error) {
	tx, err := store.connPool.Begin(ctx)

	if err != nil {
		log.Info().Msg("Can not begin transaction")
	}
	defer tx.Rollback(ctx)
	qtx := store.WithTx(tx)
	user, err := qtx.CreateAccount(ctx, arg.Params)

	if err != nil {
		return user, err
	}
	err = arg.AfterFunction(user.Email)
	if err != nil {
		return user, err
	}
	tx.Commit(ctx)
	return user, nil
}
