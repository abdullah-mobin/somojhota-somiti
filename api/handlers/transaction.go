package handlers

import (
	"context"
	"strings"

	"github.com/abdullah-mobin/somojhota-somiti/api/dtos"
	"github.com/abdullah-mobin/somojhota-somiti/api/repository"
	"github.com/abdullah-mobin/somojhota-somiti/api/response"
	"github.com/abdullah-mobin/somojhota-somiti/utils"
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
//	@Summary		Get transactions by filter
//	@Description	retrieve a transactions by filter like business id, date etc
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id			query	string	false	"transaction id"
//	@Param			business_id	query	string	false	"business id"
//	@Param			date		query	string	false	"date range in yyyy-mm-dd"
//	@Router			/transaction/ [get]
func GetTransactions(c *fiber.Ctx) error {
	queries := c.Queries()
	eligibleFilters := []string{"id", "business_id", "date"}
	objectIDFields := []string{"business_id", "id"}
	filters, err := utils.ParseFilters(queries, eligibleFilters, objectIDFields)
	if err != nil {
		return response.BadRequestException(c, "Invalid filter parameters", []string{err.Error()})
	}

	transactions, err := repository.NewTransactionRepository().GetTransactionsByFilter(context.Background(), filters)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to get transactions", err.Error())
	}

	return response.Ok(c, "Transactions Retrieved Successfully", transactions)
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
