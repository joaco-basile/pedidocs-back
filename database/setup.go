package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func GetCollection(collectionName string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb+srv://joacobasile:92HAcKTbzhcua2Sr@pedidocs.dybqk3p.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("pedidosPro").Collection(collectionName)

	return collection
}
