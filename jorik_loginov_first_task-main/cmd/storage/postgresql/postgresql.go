package postgresqlstorage

import (
	"context"
	"github.com/akayo16/jorik_loginov_first_task/config"
	"github.com/akayo16/jorik_loginov_first_task/pkg/storage/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var (
	pgxPool *pgxpool.Pool
)

func initPostgresql(config config.PostgresqlConfig) {

	pool, err := postgresql.NewClient(context.Background(), postgresql.PostgresqlConnConfig{
		Database: config.Database,
		Password: config.Password,
		User:     config.User,
		Host:     config.Host,
		Port:     config.Port,
	}, 15)
	if err != nil {
		log.Fatal("DATABASE DISCONNECT")
	}

	pgxPool = pool
}
