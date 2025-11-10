package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateNewPgxConnexionPool() (*pgxpool.Pool, error) {

	dbpool, err := pgxpool.New(context.Background(), "postgresql://postgres.bbfsckuzadzzsdymgzmj:gF7dYGWDK9tOUzCN@aws-1-eu-west-1.pooler.supabase.com:5432/postgres")
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
