package routers

import (
	"github.com/abdullah-mobin/somojhota-somiti/api/handlers"
	"github.com/abdullah-mobin/somojhota-somiti/middlewares"
	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(route fiber.Router) {
	route.Post("/new", middlewares.IsAuthenticated, handlers.CreateTransaction)
	route.Get("/:id", middlewares.IsAuthenticated, handlers.GetTransaction)
	route.Get("/", middlewares.IsAuthenticated, handlers.GetTransactions)
	route.Put("/:id", middlewares.IsAuthenticated, handlers.UpdateTransaction)
}
