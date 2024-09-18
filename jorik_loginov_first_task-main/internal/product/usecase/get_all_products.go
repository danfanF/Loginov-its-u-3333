package productusecase

import (
	"context"
	customerror "github.com/akayo16/custom-error"
	productmodel "github.com/akayo16/jorik_loginov_first_task/internal/product/model"
	"net/http"
	"strconv"
)

func All(ctx context.Context) ([]productmodel.Product, *customerror.CustomError) {
	const op = "internal.product.usecase.get_all_product.All"

	products, errGetAllProducts := postgresqlRepository.All(ctx)
	if errGetAllProducts != nil {
		return nil, errGetAllProducts.SupplementDevMessageAndChangeCode(op+":", strconv.Itoa(http.StatusInternalServerError))
	}

	return products, nil
}
