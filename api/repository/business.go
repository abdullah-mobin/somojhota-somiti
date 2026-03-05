package repository

import (
	"context"
	"errors"
	"time"

	"github.com/abdullah-mobin/somojhota-somiti/api/dtos"
	"github.com/abdullah-mobin/somojhota-somiti/api/models"
	"github.com/abdullah-mobin/somojhota-somiti/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BusinessRepository struct {
	businessCollection *mongo.Collection
}

func NewBusinessRepository() *BusinessRepository {
	return &BusinessRepository{
		businessCollection: database.DbCollections.BusinessCollection,
	}
}

func (r *BusinessRepository) CreateNewBusiness(ctx context.Context, dto *dtos.RegisterBusinessDTO) error {
	phoneAvailable, err := r.IsPhoneNumberAvailable(ctx, dto.PhoneNumber)
	if err != nil {
		return err
	}
	if !phoneAvailable {
		return errors.New("phone number already taken")
	}

	business := &models.Business{
		Name:        dto.Name,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		Status:      models.BusinessStatusActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = r.businessCollection.InsertOne(ctx, business)
	if err != nil {
		return err
	}
	return nil
}

func (r *BusinessRepository) IsPhoneNumberAvailable(ctx context.Context, phone_number string) (bool, error) {
	var user models.User
	err := r.businessCollection.FindOne(ctx, bson.M{"phone_number": phone_number}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
	if user.ID != primitive.NilObjectID {
		return false, nil
	}
	return true, nil
}

func (r *BusinessRepository) GetBusinessByID(ctx context.Context, businessID string) (*models.Business, error) {
	objID, err := primitive.ObjectIDFromHex(businessID)
	if err != nil {
		return nil, err
	}
	var business models.Business
	err = r.businessCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&business)
	if err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *BusinessRepository) UpdateBusinessByID(ctx context.Context, businessID string, update bson.M) error {
	objID, err := primitive.ObjectIDFromHex(businessID)
	if err != nil {
		return err
	}

	var updateDoc bson.M
	for k := range update {
		if len(k) > 0 && k[0] == '$' {
			updateDoc = update
			break
		}
	}
	if updateDoc == nil {
		updateDoc = bson.M{"$set": update}
	}

	_, err = r.businessCollection.UpdateOne(ctx, bson.M{"_id": objID}, updateDoc)
	return err
}
