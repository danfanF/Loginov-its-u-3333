package productmodel

import (
	"context"
	customerror "github.com/akayo16/custom-error"
	"time"
)

type Product struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Cost      int        `json:"cost"`
	Amount    int        `json:"amount"`
	UpdatedAt *time.Time `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
}

type ProductPostgresqlRepository interface {
	All(ctx context.Context) ([]Product, *customerror.CustomError)
}
