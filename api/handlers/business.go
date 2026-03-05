package handlers

import (
	"context"
	"strings"

	"github.com/abdullah-mobin/somojhota-somiti/api/dtos"
	"github.com/abdullah-mobin/somojhota-somiti/api/repository"
	"github.com/abdullah-mobin/somojhota-somiti/api/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// Business godoc
//
//	@Summary		Register new somiti
//	@Description	create a new somiti
//	@Tags			Business
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			register	body	dtos.RegisterBusinessDTO	true	"register business dto"
//	@Router			/business/new [post]
func NewBusiness(c *fiber.Ctx) error {
	var self dtos.RegisterBusinessDTO
	c.BodyParser(&self)
	err := self.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}

	err = repository.NewBusinessRepository().CreateNewBusiness(context.Background(), &self)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to create user", err.Error())
	}
	return response.Created(c, "Business Registered Successfully", nil)
}

// Business godoc
//
//	@Summary		get a somiti
//	@Description	get a somiti by its id
//	@Tags			Business
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	string	true	"business id"
//	@Router			/business/{id} [get]
func GetBusiness(c *fiber.Ctx) error {
	id := c.Params("id")

	business, err := repository.NewBusinessRepository().GetBusinessByID(context.Background(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, err.Error(), nil)
		}
		return response.InternalServerErrorException(c, "Failed to get business", err.Error())
	}
	return response.Ok(c, "Business Get Successfully", business)
}
