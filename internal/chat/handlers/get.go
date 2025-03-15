package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/yexus-api/internal/chat/services"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

func GetThreadDetails(c echo.Context) error {
	ctx := c.Request().Context()

	var body services.GetThreadDetailsParams
	if err := utils.BindAndValidate(c, &body); err != nil {
		return err
	}

	details, err := services.GetThreadDetails(ctx, body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, details)

}

func GetThreadList(c echo.Context) error {
	ctx := c.Request().Context()

	var body services.GetThreadListParams
	if err := utils.BindAndValidate(c, &body); err != nil {
		return err
	}

	res, err := services.GetThreadList(ctx, body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, utils.NewListResult(res.Threads, res.Count))
}
