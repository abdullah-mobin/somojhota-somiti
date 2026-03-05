package handlers

import (
	"context"
	"net/mail"
	"strings"

	"github.com/abdullah-mobin/somojhota-somiti/api/dtos"
	"github.com/abdullah-mobin/somojhota-somiti/api/repository"
	"github.com/abdullah-mobin/somojhota-somiti/api/response"
	"github.com/abdullah-mobin/somojhota-somiti/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Register godoc
//
//	@Summary		Register new user
//	@Description	create a new user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			register	body	dtos.RegisterUserDTO	true	"register user dto"
//	@Router			/auth/register [post]
func Register(c *fiber.Ctx) error {
	var self dtos.RegisterUserDTO
	c.BodyParser(&self)
	err := self.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}

	if !utils.IsValidBDPhoneNumber(self.PhoneNumber) {
		return response.ValidationException(c, "Invalid phone number", "Phone number is not a valid Bangladeshi phone number")
	}
	_, err = mail.ParseAddress(self.Email)
	if err != nil {
		return response.ValidationException(c, "Invalid email address", "Email address is not valid")
	}

	access, refresh, err := repository.NewUserRepository().CreateNewUser(context.Background(), &self)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to create user", err.Error())
	}
	return response.Custom(c, 200, "User Registered Successfully",
		fiber.Map{
			"access_token":  *access,
			"refresh_token": *refresh,
		},
	)
}

// Login godoc
//
//	@Summary		Login to somojhota-somiti
//	@Description	login to user account
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			login	body	dtos.LoginDTO	true	"user login dto"
//	@Router			/auth/login [post]
func Login(c *fiber.Ctx) error {
	var self dtos.LoginDTO
	c.BodyParser(&self)
	err := self.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}
	if !utils.IsValidBDPhoneNumber(self.PhoneNumber) {
		return response.ValidationException(c, "Invalid phone number", "Phone number is not a valid Bangladeshi phone number")
	}

	repo := repository.NewAuthRepository()
	credential, err := repo.GetCredentialByPhoneNumber(context.Background(), self.PhoneNumber)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to retrieve user", err.Error())
	}
	if credential == nil {
		return response.NotFoundException(c, "User not found", nil)
	}
	if !utils.ComparePassword(credential.Password, self.Password) {
		return response.UnauthorizedException(c, "Invalid credentials", nil)
	}
	tokenPayload := utils.TokenPayload{
		UserID: credential.UserID.Hex(),
	}
	accessToken, err := utils.GenerateAccessToken(tokenPayload)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to generate access token", err.Error())
	}
	refreshToken, err := utils.GenerateRefreshToken(tokenPayload)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to generate refresh token", err.Error())
	}
	err = repo.UpdateCredential(context.Background(), credential.ID.Hex(), bson.M{"refresh_token": refreshToken})
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to save refresh token", err.Error())
	}
	return response.Ok(c, "User Login successfully", fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Logout godoc
//
//	@Summary		logout from somojhota-somiti
//	@Description	logout from user account
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Router			/auth/logout [post]
func Logout(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	repo := repository.NewAuthRepository()
	err := repo.InvalidateRefreshTokens(context.Background(), userId.(string))
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to logout user", err.Error())
	}
	return response.Ok(c, "User Logout successfully", nil)
}

// Refresh token godoc
//
//	@Summary		Refresh token
//	@Description	refresh users access token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			refresh	body	dtos.RefreshTokenDTO	true	"refresh dto"
//	@Router			/auth/refresh [post]
func RefreshToken(c *fiber.Ctx) error {
	var self dtos.RefreshTokenDTO
	c.BodyParser(&self)
	err := self.Validate()
	if err != nil {
		errorsArr := strings.Split(err.Error(), ";")
		return response.ValidationException(c, "Invalid request", errorsArr)
	}
	repo := repository.NewAuthRepository()
	err, credential := repo.FindCredentialUsingRefreshToken(context.Background(), self.RefreshToken)

	if err != nil {
		return response.InternalServerErrorException(c, "Failed to retrieve credential", err.Error())
	}
	if credential == nil {
		return response.NotFoundException(c, "User not found", nil)
	}
	tokenPayload := utils.TokenPayload{
		UserID: credential.UserID.Hex(),
	}
	accessToken, err := utils.GenerateAccessToken(tokenPayload)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to generate access token", err.Error())
	}
	refreshToken, err := utils.GenerateRefreshToken(tokenPayload)
	if err != nil {
		return response.InternalServerErrorException(c, "Failed to generate refresh token", err.Error())
	}
	return response.Ok(c, "Token Refreshed successfully", fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
