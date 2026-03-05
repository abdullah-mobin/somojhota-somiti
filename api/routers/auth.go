package routers

import (
	"github.com/abdullah-mobin/somojhota-somiti/api/handlers"
	"github.com/abdullah-mobin/somojhota-somiti/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(route fiber.Router) {
	route.Post("/register", handlers.Register)
	route.Post("/login", handlers.Login)
	route.Post("/refresh", handlers.RefreshToken)
	route.Post("/logout", middlewares.IsAuthenticated, handlers.Logout)
}
