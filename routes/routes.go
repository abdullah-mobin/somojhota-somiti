package routes

import (
	"github.com/abdullah-mobin/somojhota-somiti/api/routers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	routers.AuthRoutes(api.Group("/auth"))
	routers.UserRoutes(api.Group("/user"))
	routers.BusinessRoutes(api.Group("/business"))
	routers.TransactionRoutes(api.Group("/transaction"))
	routers.BalanceSheetRoutes(api.Group("/balance-sheet"))
	routers.BalanceTransactionRoutes(api.Group("/balance-transaction"))
	routers.EarnRoutes(api.Group("/earn"))
	routers.CostRoutes(api.Group("/cost"))
}
