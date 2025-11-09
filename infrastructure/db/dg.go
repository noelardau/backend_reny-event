package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateNewPgxConnexionPool() (*pgxpool.Pool, error) {

	dbpool, err := pgxpool.New(context.Background(), "postgresql://postgres:BoissonXXLenergy261001..@localhost:5432/reny_event")
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
