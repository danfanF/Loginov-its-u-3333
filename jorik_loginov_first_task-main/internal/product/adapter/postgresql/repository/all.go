package productpostgresql

import (
	"context"
	"fmt"
	customerror "github.com/akayo16/custom-error"
	productmodel "github.com/akayo16/jorik_loginov_first_task/internal/product/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) All(ctx context.Context) ([]productmodel.Product, *customerror.CustomError) {
	const op = "internal.product.adapter.postgresql.repository.all.ALL"

	products := make([]productmodel.Product, 0)

	q := `
		SELECT *
		FROM product
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {

		if customerror.As(err, pgx.ErrNoRows) {
			return products, nil
		}

		return nil, customerror.NewCustomError(
			err,
			"Error Get All Products",
			fmt.Sprintf("%s: Error Get All Products: %s", op, err.Error()),
			"",
			op,
		)
	}

	for rows.Next() {
		var product productmodel.Product

		if err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Cost,
			&product.Amount,
			&product.UpdatedAt,
			&product.CreatedAt,
		); err != nil {
			return nil, customerror.NewCustomError(
				err,
				"Error Get All Category",
				fmt.Sprintf("%s Error Get All Category: %s", op, err.Error()),
				"",
				op,
			)
		}

		products = append(products, product)
	}

	return products, nil
}
