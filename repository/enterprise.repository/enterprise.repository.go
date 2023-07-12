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

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&Etrps)
	if err != nil {
		return Etrps, err
	}

	return Etrps, err
}

func AddOrder(newOrder m.Order, id primitive.ObjectID) error {
	//Empresa con datos para cambiar
	var EtrpsToChange = new(m.Enterprise)

	//Se almacenan los datos en la variable anterior
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(EtrpsToChange)
	if err != nil {
		return err
	}
	//Se agrega las nuevas ordenes
	EtrpsToChange.Orders = append(EtrpsToChange.Orders, newOrder)
	fmt.Println(EtrpsToChange.Orders)
	//Se actualiza la empresa con los nuevos cambios en la DB
	_, err = collection.UpdateByID(ctx, bson.M{"_id": id}, bson.D{{Key: "$set", Value: bson.D{{Key: "orders", Value: EtrpsToChange.Orders}}}})
	if err != nil {
		return err
	}

	return nil
}
