package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/yexus-api/internal/orders/services"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

func CancelAll(c echo.Context) error {
	var (
		body services.CancelAllOrdersBody
		err  error
		ctx  = c.Request().Context()
	)

	if err := utils.BindAndValidate(c, &body); err != nil {
		return err
	}

	err = services.CancelAllOrders(ctx, body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
