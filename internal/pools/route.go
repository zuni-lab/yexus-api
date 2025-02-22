package pools

import (
	"github.com/labstack/echo/v4"
	"github.com/zuni-lab/dexon-service/internal/pools/handlers"
)

func Route(g *echo.Group, path string) {
	// TODO: add middleware here

	middleware := echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TODO: add middleware here
			return next(c)
		}
	})

	pricesGroup := g.Group(path, middleware)

	pricesGroup.POST("/market", handlers.GetMarketData)
}
