package db

import (
	"context"
	"pl-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ViewProducts(client *mongo.Client) (*[]models.Product, error) {
	coll := client.Database("pl-db").Collection("products")
	// Filtro vacio que significa que busque todos
	filter := bson.D{{}}
	// Busca en la base de datos
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	// Una vez buscado lo guardamos en una variable
	var results []models.Product
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	// Finalmente retornamos los resultados.
	return &results, nil
}

func DeleteProduct(client *mongo.Client, id string) (*mongo.DeleteResult, error) {
	coll := client.Database("pl-db").Collection("products")

	// El filtro para borrar segun la ID especificada
	filter := bson.M{
		"id": id,
	}

	result, err := coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	// Retornamos que producto se elimino
	return result, nil
}

func NewProduct(client *mongo.Client, name string, id *string, price, amount int) error {
	coll := client.Database("pl-db").Collection("products")

	product := models.Product{
		ID:     id,
		Name:   name,
		Amount: amount,
		Price:  price,
	}

	_, err := coll.InsertOne(
		context.TODO(),
		product,
	)

	if err != nil {
		return err
	}

	return nil
}
