package db

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	mongopass := os.Getenv("MONGOPASS")
	if mongopass == "" {
		return nil, errors.New("An error occurred while trying to connect to MongoDB, please check the .env file.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongopass))
	if err != nil {
		return nil, err
	}

	return client, nil
}
