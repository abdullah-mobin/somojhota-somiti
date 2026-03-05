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

// Transaction godoc
//
//	@Summary		Create a new transaction
//	@Description	create a new transaction for a business
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			transaction	body	dtos.CreateTransactionDTO	true	"create transaction dto"
//	@Router			/transaction/new [post]
func CreateTransaction(c *fiber.Ctx) error {
	var dto dtos.CreateTransactionDTO
	if err := c.BodyParser(&dto); err != nil {
		return response.BadRequestException(c, "Invalid request payload", err.Error())
	}

	err := dto.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}

	transactionID, err := repository.NewTransactionRepository().CreateTransaction(context.Background(), &dto)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to create transaction", err.Error())
	}

	return response.Created(c, "Transaction Created Successfully", fiber.Map{"id": *transactionID})
}

// Transaction godoc
//
//	@Summary		Get a transaction by ID
//	@Description	retrieve a transaction by its ID
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	string	true	"transaction id"
//	@Router			/transaction/{id} [get]
func GetTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	transaction, err := repository.NewTransactionRepository().GetTransactionByID(context.Background(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, "Transaction not found", nil)
		}
		return response.InternalServerErrorException(c, "Failed to get transaction", err.Error())
	}

	return response.Ok(c, "Transaction Retrieved Successfully", transaction)
}

// Transaction godoc
//
//	@Summary		Update a transaction
//	@Description	update an existing transaction's amount or balance
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id			path	string						true	"transaction id"
//	@Param			transaction	body	dtos.UpdateTransactionDTO	true	"update transaction dto"
//	@Router			/transaction/{id} [put]
func UpdateTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	var dto dtos.UpdateTransactionDTO
	if err := c.BodyParser(&dto); err != nil {
		return response.BadRequestException(c, "Invalid request payload", err.Error())
	}
	err := dto.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}

	err = repository.NewTransactionRepository().UpdateTransactionByID(context.Background(), id, &dto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, err.Error(), nil)
		}
		return response.InternalServerErrorException(c, "Failed to update transaction", err.Error())
	}

	return response.Ok(c, "Transaction Updated Successfully", nil)
}
