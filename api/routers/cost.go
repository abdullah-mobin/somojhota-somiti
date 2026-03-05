package routers

import (
	"github.com/abdullah-mobin/somojhota-somiti/api/handlers"
	"github.com/abdullah-mobin/somojhota-somiti/middlewares"
	"github.com/gofiber/fiber/v2"
)

func CostRoutes(route fiber.Router) {
	route.Post("/new", middlewares.IsAuthenticated, handlers.CreateCost)
	route.Get("/:id", middlewares.IsAuthenticated, handlers.GetCost)
	route.Put("/:id", middlewares.IsAuthenticated, handlers.UpdateCost)
}
