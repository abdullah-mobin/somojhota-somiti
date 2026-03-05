package routers

import (
	"github.com/abdullah-mobin/somojhota-somiti/api/handlers"
	"github.com/abdullah-mobin/somojhota-somiti/middlewares"
	"github.com/gofiber/fiber/v2"
)

func EarnRoutes(route fiber.Router) {
	route.Post("/new", middlewares.IsAuthenticated, handlers.CreateEarn)
	route.Get("/:id", middlewares.IsAuthenticated, handlers.GetEarn)
	route.Put("/:id", middlewares.IsAuthenticated, handlers.UpdateEarn)
}
