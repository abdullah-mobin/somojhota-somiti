package handlers

import (
	"context"

	"github.com/abdullah-mobin/somojhota-somiti/api/repository"
	"github.com/abdullah-mobin/somojhota-somiti/api/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// User godoc
//
//	@Summary		get a user
//	@Description	get a user by its id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Router			/user/ [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Locals("userId").(string)

	business, err := repository.NewUserRepository().GetUserByID(context.Background(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, err.Error(), nil)
		}
		return response.InternalServerErrorException(c, "Failed to get user", err.Error())
	}
	return response.Ok(c, "User Get Successfully", business)
}
