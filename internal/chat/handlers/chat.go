package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/dexon-service/internal/chat/services"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

func ChatDex(c echo.Context) error {
	ctx := c.Request().Context()

	var body services.ChatParams
	if err := utils.BindAndValidate(c, &body); err != nil {
		return err
	}

	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	return services.ChatDex(ctx, body, c.Response().Writer)
}

