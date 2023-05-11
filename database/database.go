package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func GetCollection(collectionName string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://joacobasile:joacodev123@localhost:27017/?authSource=admin")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("aca tas")
		log.Fatal(err)
	}

	collection := client.Database("pedidosPro").Collection(collectionName)

	return collection
}
