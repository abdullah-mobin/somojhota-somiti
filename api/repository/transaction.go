package repository

import (
	"context"
	"errors"
	"time"

	"github.com/abdullah-mobin/somojhota-somiti/api/dtos"
	"github.com/abdullah-mobin/somojhota-somiti/api/models"
	"github.com/abdullah-mobin/somojhota-somiti/database"
	"github.com/abdullah-mobin/somojhota-somiti/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository struct {
	transactionCollection *mongo.Collection
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		transactionCollection: database.DbCollections.TransactionCollection,
	}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, dto *dtos.CreateTransactionDTO) (*primitive.ObjectID, error) {
	businessID, err := primitive.ObjectIDFromHex(dto.BusinessID)
	if err != nil {
		return nil, err
	}

	userID, err := primitive.ObjectIDFromHex(dto.UserID)
	if err != nil {
		return nil, err
	}

	if dto.Amount < 0 {
		return nil, errors.New("amount cannot be negative")
	}
	if dto.Balance < 0 {
		return nil, errors.New("balance cannot be negative")
	}

	if dto.Balance > dto.Amount {
		return nil, errors.New("invalid balance entry")
	}

	date, err := utils.ParseDateOnly(dto.Date)
	if err != nil {
		return nil, errors.New("invalid date format")
	}
	transaction := &models.Transaction{
		BusinessID: businessID,
		UserID:     userID,
		Amount:     dto.Amount,
		Balance:    dto.Balance,
		Date:       date,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	result, err := r.transactionCollection.InsertOne(ctx, transaction)
	if err != nil {
		return nil, err
	}

	m, y := utils.ParseMonthYear(dto.Date)

	balanceSheet := &models.BalanceSheet{
		BusinessID: businessID,
		Total:      dto.Balance,
		Month:      m,
		Year:       y,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	sheetRepo := NewBalanceSheetRepository()
	existingSheet, ok := sheetRepo.GetExistingSheet(context.Background(), dto.BusinessID, m, y)
	if !ok {
		if err := sheetRepo.CreateBalanceSheet(context.Background(), balanceSheet); err != nil {
			return nil, errors.New("failed to create balance sheet")
		}
	} else {
		if err := sheetRepo.Increment(context.Background(), existingSheet, dto.Balance); err != nil {
			return nil, errors.New("failed to update balance sheet")
		}
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	return &insertedID, nil
}

func (r *TransactionRepository) GetTransactionByID(ctx context.Context, transactionID string) (*models.Transaction, error) {
	objID, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		return nil, err
	}

	var transaction models.Transaction
	err = r.transactionCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) UpdateTransactionByID(ctx context.Context, transactionID string, dto *dtos.UpdateTransactionDTO) error {
	objID, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		return err
	}

	var dateStr string
	var existingTransaction models.Transaction
	err = r.transactionCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&existingTransaction)
	if err != nil {
		return err
	}
	dateStr = existingTransaction.Date.String()

	updateData := bson.M{
		"updated_at": time.Now(),
	}

	var date time.Time
	if dto.Date != nil {
		date, err = utils.ParseDateOnly(*dto.Date)
		if err != nil {
			return errors.New("invalid date format")
		}
		updateData["date"] = date
	}
	if dto.Amount != nil {
		if *dto.Amount < 0 {
			return errors.New("amount cannot be negative")
		}
		updateData["amount"] = *dto.Amount
	}
	if dto.Balance != nil {
		if *dto.Balance < 0 {
			return errors.New("balance cannot be negative")
		}
		updateData["balance"] = *dto.Balance
		m, y := utils.ParseMonthYear(dateStr)
		sheetRepo := NewBalanceSheetRepository()
		existingSheet, _ := sheetRepo.GetExistingSheet(context.Background(), existingTransaction.BusinessID.Hex(), m, y)
		sheetRepo.Decrement(context.Background(), existingSheet, existingSheet.Total)
		sheetRepo.Increment(context.Background(), existingSheet, *dto.Balance)
	}

	_, err = r.transactionCollection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateData},
	)

	return err
}
