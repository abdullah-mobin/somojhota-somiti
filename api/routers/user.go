package routers

import (
	"github.com/abdullah-mobin/somojhota-somiti/api/handlers"
	"github.com/abdullah-mobin/somojhota-somiti/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(route fiber.Router) {
	route.Get("/", middlewares.IsAuthenticated, handlers.GetUser)
	route.Get("/:id", middlewares.IsAuthenticated, handlers.GetUserByID)
}
