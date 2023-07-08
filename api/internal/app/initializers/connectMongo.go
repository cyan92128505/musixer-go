package initializers

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	migrate "github.com/xakep666/mongo-migrate"
)

var DB *mongo.Database

func ConnectDB(config *Config) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBUrl))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")
	DB = client.Database(config.DBName)
}

func Migration(db *mongo.Database) (*mongo.Database, error) {
	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		return nil, err
	}
	return db, nil
}
