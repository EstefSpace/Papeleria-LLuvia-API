package db

/*
Validación de ApiKeys y Creación de los mismos.
Aun no aplico esto a la api.
*/

import (
	"context"
	"pl-api/models"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateApiKey(client *mongo.Client, webClient string) error {

	// Uso de la libreria gonanoid para generar una apikey de 36 caracteres con las letras del abecedario y los numeros del 0 al 9
	apikey, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 36)

	if err != nil {
		return err
	}

	coll := client.Database("pl-db").Collection("apikeys")

	clientAPI := models.ClientAPI{
		ApiKey:    apikey,
		WebClient: webClient,
	}

	_, err = coll.InsertOne(
		context.TODO(),
		clientAPI,
	)

	if err != nil {
		return err
	}

	return nil
}
