package user_repository

import (
	"context"
	"time"

	"github.com/joaco-basile/pedidocs-back/database"
	m "github.com/joaco-basile/pedidocs-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func LoginUser(nameOrEmail string, password string) (m.User, error) {
	var user m.User

	err := collection.FindOne(ctx, bson.D{{Key: "name", Value: nameOrEmail}}).Decode(&user)
	if err != nil {
		err = collection.FindOne(ctx, bson.D{{Key: "email", Value: nameOrEmail}}).Decode(&user)
		if err != nil {
			return user, err
		}
	}

	return user, nil

}

func RegisterUser(user m.User) error {
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user m.User, id primitive.ObjectID) error {

	var err error

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"password":   user.Password,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id primitive.ObjectID) error {

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}

	return nil
}
