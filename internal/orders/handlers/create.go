package handlers

import (
	"net/http"

	"github.com/zuni-lab/yexus-api/internal/orders/services"

	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

func Create(c echo.Context) error {
	var (
		body services.CreateOrderBody
		ctx  = c.Request().Context()
	)

	if err := utils.BindAndValidate(c, &body); err != nil {
		return err
	}

	order, err := services.CreateOrder(ctx, body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, order)
}
