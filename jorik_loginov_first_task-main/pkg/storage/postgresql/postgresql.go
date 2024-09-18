package postgresql

import (
	"context"
	"fmt"
	"github.com/akayo16/jorik_loginov_first_task/pkg/lib/logger/sl"
	"github.com/akayo16/jorik_loginov_first_task/pkg/lib/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type PostgresqlConnConfig struct {
	Database string
	Password string
	User     string
	Host     string
	Port     string
}

func NewClient(
	ctx context.Context,
	cfg PostgresqlConnConfig,
	maxAttempts int,
) (pool *pgxpool.Pool, err error) {
	const op = "pkg.storage.postgresql.newClient"

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	//dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal(fmt.Sprintf("ERROR DO WITH TRIES POSTGRESQL: %s", op), sl.Err(err))
	}

	return pool, nil
}
