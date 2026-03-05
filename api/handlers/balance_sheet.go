package handlers

import (
	"context"

	"github.com/abdullah-mobin/somojhota-somiti/api/repository"
	"github.com/abdullah-mobin/somojhota-somiti/api/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// Balance-Sheet godoc
//
//	@Summary		Get a balance sheet
//	@Description	retrieve a balance sheet by business ID
//	@Tags			Balance-sheet
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			business_id	path	string	true	"business id"
//	@Router			/balance-sheet/{business_id} [get]
func GetBalanceSheet(c *fiber.Ctx) error {
	bID := c.Params("business_id")

	transaction, err := repository.NewBalanceSheetRepository().GetBalanceSheetByBsuinessID(context.Background(), bID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundException(c, "balance sheet not found", nil)
		}
		return response.InternalServerErrorException(c, "Failed to get balance sheet", err.Error())
	}

	return response.Ok(c, "Balance Sheet Retrieved Successfully", transaction)
}
