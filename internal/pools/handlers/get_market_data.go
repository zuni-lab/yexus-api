package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/dexon-service/internal/pools/services"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

func GetMarketData(c echo.Context) error {
	ctx := c.Request().Context()

	var body services.GetMarketDataParams
	if err := utils.BindAndValidate(c, &body); err != nil {
		return err
	}

	candlestick, err := services.GetMarketData(ctx, body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, candlestick)
}
