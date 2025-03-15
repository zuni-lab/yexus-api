package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zuni-lab/yexus-api/internal/orders/services"

	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

func Get(c echo.Context) error {
	var (
		query services.GetOrderByIDQuery
		err   error
		ctx   = c.Request().Context()
	)
	if err := utils.BindAndValidate(c, &query); err != nil {
		return err
	}

	query.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid id"))
	}

	order, err := services.GetOrderByID(ctx, query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, order)
}
