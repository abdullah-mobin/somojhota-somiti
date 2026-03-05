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

type UserRepository struct {
	userCollection       *mongo.Collection
	credentialCollection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		userCollection:       database.DbCollections.UserCollection,
		credentialCollection: database.DbCollections.CredentialCollection,
	}
}

func (r *UserRepository) CreateNewUser(ctx context.Context, dto *dtos.RegisterUserDTO) (*string, *string, error) {
	phoneAvailable, err := NewAuthRepository().IsPhoneNumberAvailable(ctx, dto.PhoneNumber)
	if err != nil {
		return nil, nil, err
	}
	if !phoneAvailable {
		return nil, nil, errors.New("phone number already taken")
	}

	user := &models.User{
		Name:        dto.Name,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		Status:      models.StatusActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbUser, err := r.userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	uId, _ := dbUser.InsertedID.(primitive.ObjectID)
	accessToken, _ := utils.GenerateAccessToken(utils.TokenPayload{UserID: uId.String()})
	refreshToken, _ := utils.GenerateRefreshToken(utils.TokenPayload{UserID: uId.String()})
	credential := &models.Credential{
		UserID:       uId,
		PhoneNumber:  dto.PhoneNumber,
		Password:     utils.HashPassword(dto.Password),
		RefreshToken: refreshToken.(string),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err = NewAuthRepository().CreateCredential(ctx, credential)
	if err != nil {
		return nil, nil, err
	}
	access := accessToken.(string)
	refresh := refreshToken.(string)
	return &access, &refresh, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	var data models.User
	err = r.userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// func (r *UserRepository) UpdateUserByID(ctx context.Context, userID string, update bson.M) error {
// 	objID, err := primitive.ObjectIDFromHex(userID)
// 	if err != nil {
// 		return err
// 	}

// 	var updateDoc bson.M
// 	for k := range update {
// 		if len(k) > 0 && k[0] == '$' {
// 			updateDoc = update
// 			break
// 		}
// 	}
// 	if updateDoc == nil {
// 		updateDoc = bson.M{"$set": update}
// 	}

// 	_, err = r.userCollection.UpdateOne(ctx, bson.M{"_id": objID}, updateDoc)
// 	_, err = r.credentialCollection.UpdateOne(ctx, bson.M{"user_id": objID}, )
// 	return err
// }
