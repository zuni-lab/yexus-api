package handlers

import (
	"net/http"

	"github.com/zuni-lab/dexon-service/internal/orders/services"

	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

func List(c echo.Context) error {
	var (
		query services.ListOrdersByWalletQuery
		ctx   = c.Request().Context()
	)
	if err := utils.BindAndValidate(c, &query); err != nil {
		return err
	}

	results, err := services.ListOrderByWallet(ctx, query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, utils.NewListResult(results.Orders, results.Count))
}
