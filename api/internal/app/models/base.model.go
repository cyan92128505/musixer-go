package models

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *BaseModel) Create(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {
	collection := db.Collection(collectionName)

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	m.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (m *BaseModel) Read(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (m *BaseModel) Update(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
	collection := db.Collection(collectionName)

	m.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *BaseModel) Delete(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

var validate = validator.New()

type SuccessResponse struct {
	Status string `json:"status"`
}

type TokenResponse struct {
	SuccessResponse
	AccessToken string `json:"access_token"`
}

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
