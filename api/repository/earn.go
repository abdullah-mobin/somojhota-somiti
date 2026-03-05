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

type EarnRepository struct {
	earnCollection *mongo.Collection
}

func NewEarnRepository() *EarnRepository {
	return &EarnRepository{
		earnCollection: database.DbCollections.EarnCollection,
	}
}

func (r *EarnRepository) CreateEarn(ctx context.Context, dto *dtos.CreateEarnDTO) (*primitive.ObjectID, error) {
	businessID, err := primitive.ObjectIDFromHex(dto.BusinessID)
	if err != nil {
		return nil, err
	}

	userID, err := primitive.ObjectIDFromHex(dto.UserID)
	if err != nil {
		return nil, err
	}

	if dto.Balance < 0 {
		return nil, errors.New("balance cannot be negative")
	}

	date, err := utils.ParseDateOnly(dto.Date)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	earn := &models.Earn{
		BusinessID:  businessID,
		UserID:      userID,
		Description: dto.Description,
		Balance:     dto.Balance,
		Date:        date,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := r.earnCollection.InsertOne(ctx, earn)
	if err != nil {
		return nil, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	return &insertedID, nil
}

func (r *EarnRepository) GetEarnByID(ctx context.Context, earnID string) (*models.Earn, error) {

	objID, err := primitive.ObjectIDFromHex(earnID)
	if err != nil {
		return nil, err
	}

	var earn models.Earn

	err = r.earnCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&earn)
	if err != nil {
		return nil, err
	}

	return &earn, nil
}

func (r *EarnRepository) UpdateEarnByID(ctx context.Context, earnID string, dto *dtos.UpdateEarnDTO) error {

	objID, err := primitive.ObjectIDFromHex(earnID)
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

	_, err = r.earnCollection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateData},
	)

	return err
}
