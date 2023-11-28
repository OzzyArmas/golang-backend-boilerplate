package db

import (
	"context"

	"github.com/BowlFinder/bowl-finder-server/sqlc_db"
	"github.com/jackc/pgx/v5"
)

func New() (*sqlc_db.Queries, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		return &sqlc_db.Queries{}, err
	}
	defer conn.Close(ctx)

	client := sqlc_db.New(conn)

	return client, nil
}
