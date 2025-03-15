package server

import (
	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/yexus-api/internal/chat"
	"github.com/zuni-lab/yexus-api/internal/health"
	"github.com/zuni-lab/yexus-api/internal/orders"
)

func setupRoute(e *echo.Echo) {
	api := e.Group("/api")
	health.Route(e, "/health")
	chat.Route(api, "/chat")
	orders.Route(api, "/orders")
}
