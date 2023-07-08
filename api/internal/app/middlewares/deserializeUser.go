package middleware

import (
	"log"
	"musixer/api/internal/app/initializers"
	"musixer/api/internal/app/models"
	"musixer/api/internal/app/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeserializeUser(c *fiber.Ctx) error {
	var access_token string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		access_token = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		access_token = c.Cookies("access_token")
	}

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := initializers.LoadConfig(".")

	tokenClaims, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "ValidateTokenError:" + err.Error()})
	}

	userID, err := initializers.RedisClient.Get(c.UserContext(), tokenClaims.TokenUuid).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Token is invalid or session has expired"})
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Fatalf("轉換 _id 時發生錯誤：%s\n", err)
	}

	filter := bson.M{"_id": objectID}
	var user models.User
	err = user.Find(c.UserContext(), initializers.DB, models.UserCollectionName, filter, &user)
	if err != nil {
		log.Fatal(err)
	}

	c.Locals("user", models.FilterUserRecord(&user))
	c.Locals("access_token_uuid", tokenClaims.TokenUuid)

	return c.Next()
}
