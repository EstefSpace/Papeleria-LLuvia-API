package db

import (
	"context"
	"pl-api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewSale(client *mongo.Client, id *string, user string, total float64, date string, products []models.Product) error {
	coll := client.Database("pl-db").Collection("sales")

	sale := models.Sale{
		ID:       id,
		User:     user,
		Total:    total,
		Date:     date,
		Products: products,
	}
	_, err := coll.InsertOne(
		context.TODO(), sale,
	)

	if err != nil {
		return err
	}

	return nil
}
