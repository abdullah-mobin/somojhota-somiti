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

// Cost godoc
//
//	@Summary		Create a new cost
//	@Description	create a new cost entry
//	@Tags			Cost
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			cost	body	dtos.CreateCostDTO	true	"create cost dto"
//	@Router			/cost/new [post]
func CreateCost(c *fiber.Ctx) error {
	var dto dtos.CreateCostDTO

	if err := c.BodyParser(&dto); err != nil {
		return response.BadRequestException(c, "Invalid request payload", err.Error())
	}

	err := dto.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}

	costID, err := repository.NewCostRepository().CreateCost(context.Background(), &dto)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to create cost", err.Error())
	}

	return response.Created(c, "Cost Created Successfully", fiber.Map{"id": *costID})
}

// Cost godoc
//
//	@Summary		Get a cost by ID
//	@Description	retrieve cost by id
//	@Tags			Cost
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	string	true	"cost id"
//	@Router			/cost/{id} [get]
func GetCost(c *fiber.Ctx) error {
	id := c.Params("id")

	cost, err := repository.NewCostRepository().GetCostByID(context.Background(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, "Cost not found", nil)
		}
		return response.InternalServerErrorException(c, "Failed to get cost", err.Error())
	}

	return response.Ok(c, "Cost Retrieved Successfully", cost)
}

// Cost godoc
//
//	@Summary		Update a cost
//	@Description	update an existing cost
//	@Tags			Cost
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path	string				true	"cost id"
//	@Param			cost	body	dtos.UpdateCostDTO	true	"update cost dto"
//	@Router			/cost/{id} [put]
func UpdateCost(c *fiber.Ctx) error {
	id := c.Params("id")

	var dto dtos.UpdateCostDTO
	if err := c.BodyParser(&dto); err != nil {
		return response.BadRequestException(c, "Invalid request payload", err.Error())
	}

	err := dto.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}

	err = repository.NewCostRepository().UpdateCostByID(context.Background(), id, &dto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, "Cost not found", nil)
		}
		return response.InternalServerErrorException(c, "Failed to update cost", err.Error())
	}

	return response.Ok(c, "Cost Updated Successfully", nil)
}
