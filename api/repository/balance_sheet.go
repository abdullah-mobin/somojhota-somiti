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
}

func NewBalanceSheetRepository() *BalanceSheetRepository {
	return &BalanceSheetRepository{
		balanceSheetCollection: database.DbCollections.BalanceSheetCollection,
	}
}

func (r *BalanceSheetRepository) GetBalanceSheetByBsuinessID(ctx context.Context, businessID string) (*models.BalanceSheet, error) {
	objID, err := primitive.ObjectIDFromHex(businessID)
	if err != nil {
		return nil, err
	}

	var data models.BalanceSheet
	err = r.balanceSheetCollection.FindOne(ctx, bson.M{"business_id": objID}).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
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
