package enterprise_repository

import (
	"context"
	"fmt"

	"github.com/joaco-basile/pedidocs-back/database"
	m "github.com/joaco-basile/pedidocs-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()
var collection = database.GetCollection("enterprise")

func CreateEnterprise(etrps m.Enterprise) error {
	_, err := collection.InsertOne(ctx, etrps)
	if err != nil {
		return err
	}
	return nil
}

func GetEnterprise(id primitive.ObjectID) (m.Enterprise, error) {
	var Etrps m.Enterprise
	fmt.Println(id)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&Etrps)
	if err != nil {
		return Etrps, err
	}

	return Etrps, nil
}

func AddOrder(newOrder m.Order, id primitive.ObjectID) error {

	filtro := bson.M{"_id": id}
	actualizacion := bson.M{"$push": bson.M{"orders": newOrder}}

	_, err := collection.UpdateOne(ctx, filtro, actualizacion)
	if err != nil {
		return err
	}

	return nil
}
