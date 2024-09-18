package httphandlersv1product

import (
	"context"
	"encoding/json"
	"fmt"
	productusecase "github.com/akayo16/jorik_loginov_first_task/internal/product/usecase"
	"log/slog"
	"net/http"
	"strconv"
)

// @Summary Get All Product
// @Tags Product
// @ID get-all-product
// @Success 200 {object} []productmodel.Product
// @Router /product/get/all [get]
func All(w http.ResponseWriter, r *http.Request) {
	const op = "All Product Endpoint"

	allProduct, customErr := productusecase.All(context.Background())
	if customErr != nil {
		statusError, err := strconv.Atoi(customErr.Code())
		if err != nil {
			slog.Error(fmt.Sprintf("%s: Error Converting CustomErr.Code() To HTTP Status Code: %s", op, err.Error()))

			statusError = http.StatusInternalServerError
		}

		http.Error(w, customErr.Message(), statusError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allProduct); err != nil {
		slog.Error(fmt.Sprintf("%s: Error Encoding Response: %s", op, err.Error()))

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
