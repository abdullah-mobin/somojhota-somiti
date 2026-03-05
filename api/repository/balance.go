package repository

import (
	"context"
	"time"

	"github.com/abdullah-mobin/somojhota-somiti/api/models"
	"github.com/abdullah-mobin/somojhota-somiti/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BalanceSheetRepository struct {
	balanceSheetCollection *mongo.Collection
	transactionCollection  *mongo.Collection
	userCollection         *mongo.Collection
}

func NewBalanceSheetRepository() *BalanceSheetRepository {
	return &BalanceSheetRepository{
		balanceSheetCollection: database.DbCollections.BalanceSheetCollection,
		transactionCollection:  database.DbCollections.TransactionCollection,
		userCollection:         database.DbCollections.UserCollection,
	}
}

type BalanceSheetSummary struct {
	Balance          float64               `json:"balance"`
	TotalTransaction float64               `json:"total_transaction"`
	BalanceSheet     []models.BalanceSheet `json:"balance_sheet"`
}

type Transactions struct {
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`
	Balance float64 `json:"balance"`
}

type BalanceTransactionSummary struct {
	Amount       float64        `json:"amount"`
	Balance      float64        `json:"balance"`
	Transactions []Transactions `json:"transactions"`
}

func (r *BalanceSheetRepository) GetBalanceSheetByBsuinessID(ctx context.Context, businessID string) (*BalanceSheetSummary, error) {
	objID, err := primitive.ObjectIDFromHex(businessID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.balanceSheetCollection.Find(ctx, bson.M{
		"business_id": objID,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var data []models.BalanceSheet
	if err := cursor.All(ctx, &data); err != nil {
		return nil, err
	}

	var totalTransaction float64
	for _, sheet := range data {
		totalTransaction += sheet.Total
	}

	return &BalanceSheetSummary{
		TotalTransaction: totalTransaction,
		BalanceSheet:     data,
	}, nil
}

func (r *BalanceSheetRepository) GetBalanceTransactionByBusinessID(ctx context.Context, businessID string) (*BalanceTransactionSummary, error) {
	objID, err := primitive.ObjectIDFromHex(businessID)
	if err != nil {
		return nil, err
	}

	user, err := NewUserRepository().GetAllUserByBusinessID(context.Background(), businessID)
	if err != nil {
		return nil, err
	}

	var summary BalanceTransactionSummary

	for _, u := range user {
		var tran Transactions
		cursor, err := r.transactionCollection.Find(ctx, bson.M{
			"business_id": objID,
			"user_id":     u.ID,
		})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		tran.Name = u.Name
		var data []models.Transaction
		if err := cursor.All(ctx, &data); err != nil {
			return nil, err
		}
		for _, t := range data {
			tran.Amount += t.Amount
			tran.Balance += t.Balance
		}
		summary.Transactions = append(summary.Transactions, tran)
		summary.Amount += tran.Amount
		summary.Balance += tran.Balance
	}

	return &summary, nil
}

func (r *BalanceSheetRepository) GetExistingSheet(ctx context.Context, businessID string, m, y int) (*models.BalanceSheet, bool) {
	objID, _ := primitive.ObjectIDFromHex(businessID)
	var data models.BalanceSheet
	err := r.balanceSheetCollection.FindOne(ctx,
		bson.M{
			"business_id": objID,
			"month":       m,
			"year":        y,
		}).Decode(&data)
	if err == mongo.ErrNoDocuments {
		return nil, false

	}
	return &data, true
}

func (r *BalanceSheetRepository) CreateBalanceSheet(ctx context.Context, balanceSheet *models.BalanceSheet) error {
	_, err := r.balanceSheetCollection.InsertOne(ctx, balanceSheet)
	if err != nil {
		return err
	}
	return nil
}

func (r *BalanceSheetRepository) Increment(ctx context.Context, sheet *models.BalanceSheet, amount float64) error {
	update := bson.M{
		"$inc": bson.M{
			"total": amount,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := r.balanceSheetCollection.UpdateOne(
		ctx,
		bson.M{"_id": sheet.ID},
		update,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *BalanceSheetRepository) Decrement(ctx context.Context, sheet *models.BalanceSheet, amount float64) error {
	update := bson.M{
		"$inc": bson.M{
			"total": -amount,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err := r.balanceSheetCollection.UpdateOne(
		ctx,
		bson.M{"_id": sheet.ID},
		update,
	)
	if err != nil {
		return err
	}

	return nil
}
