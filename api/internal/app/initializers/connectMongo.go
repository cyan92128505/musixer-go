package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("connected successfully")
	return nil
}

func ConnectDB(config *Config) {
	var err error
	mongoConnectionString := os.Getenv("DATABASE_URL")

	DB, ctx, cancel, err := connect(mongoConnectionString)
	if err != nil {
		panic(err)
	}

	defer close(DB, ctx, cancel)

	err = ping(DB, ctx)
	if err != nil {
		panic(err)
	}

	log.Println("🚀 Connected Successfully to the Database")
}
