package routers

import (
	"github.com/abdullah-mobin/somojhota-somiti/api/handlers"
	"github.com/abdullah-mobin/somojhota-somiti/middlewares"
	"github.com/gofiber/fiber/v2"
)

func BusinessRoutes(route fiber.Router) {
	route.Post("/new", middlewares.IsAuthenticated, handlers.NewBusiness)
	route.Get("/:id", middlewares.IsAuthenticated, handlers.GetBusiness)
	// route.Put("/:id", middlewares.IsAuthenticated, handlers.Logout)
}
