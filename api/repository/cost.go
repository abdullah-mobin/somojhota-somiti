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

type CostRepository struct {
	costCollection *mongo.Collection
}

func NewCostRepository() *CostRepository {
	return &CostRepository{
		costCollection: database.DbCollections.CostCollection,
	}
}

func (r *CostRepository) CreateCost(ctx context.Context, dto *dtos.CreateCostDTO) (*primitive.ObjectID, error) {
	businessID, err := primitive.ObjectIDFromHex(dto.BusinessID)
	if err != nil {
		return nil, err
	}

	userID, err := primitive.ObjectIDFromHex(dto.UserID)
	if err != nil {
		return nil, err
	}

	date, err := utils.ParseDateOnly(dto.Date)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	cost := &models.Cost{
		BusinessID:  businessID,
		UserID:      userID,
		Description: dto.Description,
		Balance:     dto.Balance,
		Date:        date,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := r.costCollection.InsertOne(ctx, cost)
	if err != nil {
		return nil, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	return &insertedID, nil
}

func (r *CostRepository) GetCostByID(ctx context.Context, costID string) (*models.Cost, error) {

	objID, err := primitive.ObjectIDFromHex(costID)
	if err != nil {
		return nil, err
	}

	var cost models.Cost

	err = r.costCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&cost)
	if err != nil {
		return nil, err
	}

	return &cost, nil
}

func (r *CostRepository) UpdateCostByID(ctx context.Context, costID string, dto *dtos.UpdateCostDTO) error {
	objID, err := primitive.ObjectIDFromHex(costID)
	if err != nil {
		return err
	}

	updateData := bson.M{
		"updated_at": time.Now(),
	}

	if dto.Description != nil {
		updateData["description"] = *dto.Description
	}

	if dto.Balance != nil {
		if *dto.Balance < 0 {
			return errors.New("balance cannot be negative")
		}
		updateData["balance"] = *dto.Balance
	}

	if dto.Date != nil {
		date, err := utils.ParseDateOnly(*dto.Date)
		if err != nil {
			return errors.New("invalid date format")
		}
		updateData["date"] = date
	}

	_, err = r.costCollection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateData},
	)

	return err
}
