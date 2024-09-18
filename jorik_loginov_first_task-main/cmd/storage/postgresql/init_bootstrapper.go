package postgresqlstorage

import (
	productusecase "github.com/akayo16/jorik_loginov_first_task/internal/product/usecase"
	"github.com/akayo16/jorik_loginov_first_task/pkg/storage/postgresql"
)

func initAllBootstrapper() {
	initProductPostgresqlRepositoryConfiguration(pgxPool)
}

func initProductPostgresqlRepositoryConfiguration(pgx postgresql.Client) {
	productusecase.ProductPostgresqlRepositoryConfiguration(pgx)
}
