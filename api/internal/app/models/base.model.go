package models

import (
	"context"
	"time"

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
