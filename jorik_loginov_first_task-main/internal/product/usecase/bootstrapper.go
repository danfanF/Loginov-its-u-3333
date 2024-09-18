package productusecase

import (
	productpostgresql "github.com/akayo16/jorik_loginov_first_task/internal/product/adapter/postgresql/repository"
	productmodel "github.com/akayo16/jorik_loginov_first_task/internal/product/model"
	"github.com/akayo16/jorik_loginov_first_task/pkg/storage/postgresql"
)

var (
	postgresqlRepository productmodel.ProductPostgresqlRepository
)

func ProductPostgresqlRepositoryConfiguration(pgx postgresql.Client) {
	postgresqlRepository = productpostgresql.NewRepository(pgx)
}
