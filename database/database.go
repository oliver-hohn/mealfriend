package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func Start(dbURL string, c context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(c, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %w", err)
	}

	return conn, nil
}
