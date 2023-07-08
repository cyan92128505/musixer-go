package models

import (
	"context"
	"musixer/api/internal/app/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollectionName = "users"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Name      *string            `json:"name" validate:"required,min=2,max=100"`
	Email     *string            `json:"email" validate:"email,required"`
	Password  *string            `json:"password" validate:"required,min=6"`
	Verified  *bool              `json:"verified"`
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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Name:      *user.Name,
		Email:     *user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
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

func (m *User) Find(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result *User) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
