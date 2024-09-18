package productpostgresql

import (
	productmodel "github.com/akayo16/jorik_loginov_first_task/internal/product/model"
	"github.com/akayo16/jorik_loginov_first_task/pkg/storage/postgresql"
)

type repository struct {
	client postgresql.Client
}

func NewRepository(client postgresql.Client) productmodel.ProductPostgresqlRepository {
	return &repository{
		client: client,
	}
}
