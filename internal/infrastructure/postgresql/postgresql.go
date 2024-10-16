package postgresql

import (
	"context"
	"fmt"
	"teams_service/internal/core/configuration"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(config *configuration.DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s", config.User, config.Pass, config.Host, config.Name)
	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
