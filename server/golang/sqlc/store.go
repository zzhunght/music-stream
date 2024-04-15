package sqlc

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier()
}
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(pool *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		connPool: pool,
		Queries:  New(pool),
	}
}
