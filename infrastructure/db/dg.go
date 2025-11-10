package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateNewPgxConnexionPool() (*pgxpool.Pool, error) {

	dbpool, err := pgxpool.New(context.Background(), "postgresql://postgres:gF7dYGWDK9tOUzCN@db.bbfsckuzadzzsdymgzmj.supabase.co:5432/postgres")
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
