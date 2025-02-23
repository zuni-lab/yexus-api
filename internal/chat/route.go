package chat

import (
	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/dexon-service/internal/chat/handlers"
)

func Route(g *echo.Group, path string) {
	chatGroup := g.Group(path)
	chatGroup.POST("/dex", handlers.ChatDex)
}
