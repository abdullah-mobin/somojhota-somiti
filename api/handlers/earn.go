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

// Earn godoc
//
//	@Summary		Create a new earn
//	@Description	create a new earn entry
//	@Tags			Earn
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			earn	body	dtos.CreateEarnDTO	true	"create earn dto"
//	@Router			/earn/new [post]
func CreateEarn(c *fiber.Ctx) error {
	var dto dtos.CreateEarnDTO

	if err := c.BodyParser(&dto); err != nil {
		return response.BadRequestException(c, "Invalid request payload", err.Error())
	}

	err := dto.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}

	earnID, err := repository.NewEarnRepository().CreateEarn(context.Background(), &dto)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to create earn", err.Error())
	}

	return response.Created(c, "Earn Created Successfully", fiber.Map{"id": *earnID})
}

// Earn godoc
//
//	@Summary		Get a earn by ID
//	@Description	retrieve earn by id
//	@Tags			Earn
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	string	true	"earn id"
//	@Router			/earn/{id} [get]
func GetEarn(c *fiber.Ctx) error {
	id := c.Params("id")

	earn, err := repository.NewEarnRepository().GetEarnByID(context.Background(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, "Earn not found", nil)
		}
		return response.InternalServerErrorException(c, "Failed to get earn", err.Error())
	}

	return response.Ok(c, "Earn Retrieved Successfully", earn)
}

// Earn godoc
//
//	@Summary		Update a earn
//	@Description	update an existing earn
//	@Tags			Earn
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path	string				true	"earn id"
//	@Param			earn	body	dtos.UpdateEarnDTO	true	"update earn dto"
//	@Router			/earn/{id} [put]
func UpdateEarn(c *fiber.Ctx) error {
	id := c.Params("id")

	var dto dtos.UpdateEarnDTO
	if err := c.BodyParser(&dto); err != nil {
		return response.BadRequestException(c, "Invalid request payload", err.Error())
	}

	err := dto.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}

	err = repository.NewEarnRepository().UpdateEarnByID(context.Background(), id, &dto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, "Earn not found", nil)
		}
		return response.InternalServerErrorException(c, "Failed to update earn", err.Error())
	}

	return response.Ok(c, "Earn Updated Successfully", nil)
}
