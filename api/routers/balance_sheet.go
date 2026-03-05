package routers

import (
	"github.com/abdullah-mobin/somojhota-somiti/api/handlers"
	"github.com/abdullah-mobin/somojhota-somiti/middlewares"
	"github.com/gofiber/fiber/v2"
)

func BalanceSheetRoutes(route fiber.Router) {
	route.Get("/:business_id", middlewares.IsAuthenticated, handlers.GetBalanceSheet)
}
