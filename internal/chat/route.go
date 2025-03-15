package chat

import (
	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/yexus-api/internal/chat/handlers"
)

func Route(g *echo.Group, path string) {
	chatGroup := g.Group(path)
	chatGroup.POST("/dex", handlers.ChatDex)
	chatGroup.POST("/dex/thread", handlers.GetThreadDetails)
	chatGroup.POST("/dex/thread/list", handlers.GetThreadList)
}
