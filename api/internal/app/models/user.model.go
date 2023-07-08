package models

import (
	"context"
	"musixer/api/internal/app/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollectionName = "users"

type User struct {
	BaseModel
	Name     *string `json:"name" validate:"required,min=2,max=100"`
	Email    *string `json:"email" validate:"email,required"`
	Password *string `json:"password" validate:"required,min=6"`
	Role     *string `json:"role"`
	Provider *string `json:"provider"`
	Verified *bool   `json:"verified"`
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Name:      *user.Name,
		Email:     *user.Email,
		Role:      *user.Role,
		Provider:  *user.Provider,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (user *User) Register(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(UserCollectionName)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	filter := bson.M{"email": user.Email}
	existingUser := collection.FindOne(context.Background(), filter)
	if existingUser.Err() == nil {
		return utils.ErrEmailExists
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
