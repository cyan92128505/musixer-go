package controllers

import (
	"musixer/api/internal/app/initializers"
	"musixer/api/internal/app/models"
	"musixer/api/internal/app/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// @Summary SignUp User
// @Description Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param payload body models.SignUpInput true "User data"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /api/auth/signup [post]
func SignUpUser(c *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	lowerEmail := strings.ToLower(payload.Email)
	rawPassword := string(hashedPassword)

	newUser := models.User{
		Name:     &payload.Name,
		Email:    &lowerEmail,
		Password: &rawPassword,
	}

	err = newUser.Register(c.UserContext(), initializers.DB)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "ERROR:" + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
}

// @Summary Sign In User
// @Description Sign in with email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param payload body models.SignInInput true "User credentials"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /api/auth/signin [post]
func SignInUser(c *fiber.Ctx) error {
	var payload *models.SignInInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "BodyParser " + err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	message := "Invalid email or password"

	var foundUser models.User

	config, _ := initializers.LoadConfig(".")
	filter := bson.M{"email": strings.ToLower(payload.Email)}
	err := foundUser.Read(c.UserContext(), initializers.DB, models.UserCollectionName, filter, &foundUser)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": "DB Error " + err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(*foundUser.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": "Password Error " + message})
	}

	accessTokenDetails, err := utils.CreateToken(foundUser.ID.String(), config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refreshTokenDetails, err := utils.CreateToken(foundUser.ID.String(), config.RefreshTokenExpiresIn, config.RefreshTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()

	errAccess := initializers.RedisClient.Set(c.UserContext(), accessTokenDetails.TokenUuid, foundUser.ID.String(), time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": errAccess.Error()})
	}

	errRefresh := initializers.RedisClient.Set(c.UserContext(), refreshTokenDetails.TokenUuid, foundUser.ID.String(), time.Unix(*refreshTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": errRefresh.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.RefreshTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token})
}

// @Summary Refresh Access Token
// @Description Refresh the access token using a refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /api/auth/refresh [get]
func RefreshAccessToken(c *fiber.Ctx) error {
	message := "could not refresh access token"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	config, _ := initializers.LoadConfig(".")
	tokenClaims, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	userID, err := initializers.RedisClient.Get(c.UserContext(), tokenClaims.TokenUuid).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	var user models.User

	objectID, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": objectID}

	err = user.Read(c.UserContext(), initializers.DB, models.UserCollectionName, filter, user)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	accessTokenDetails, err := utils.CreateToken(user.ID.String(), config.AccessTokenExpiresIn, config.AccessTokenPrivateKey)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()

	errAccess := initializers.RedisClient.Set(c.UserContext(), accessTokenDetails.TokenUuid, user.ID.String(), time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": errAccess.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   config.AccessTokenMaxAge * 60,
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token})
}

// @Description Logout a user by invalidating access and refresh tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /api/auth/logout [get]
func LogoutUser(c *fiber.Ctx) error {
	message := "Token is invalid or session has expired"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	config, _ := initializers.LoadConfig(".")

	tokenClaims, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	access_token_uuid := c.Locals("access_token_uuid").(string)
	_, err = initializers.RedisClient.Del(c.UserContext(), tokenClaims.TokenUuid, access_token_uuid).Result()
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
