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

func LoginUser(name string, email string, password string) (m.User, error) {
	var user m.User

	if name != "" {
		err := collection.FindOne(ctx, bson.D{{"name", name}}).Decode(&user)
		if err != nil {
			return user, err
		}
	}

	if email != "" {
		err := collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
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

func UpdateUser(user m.User, userId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func DseleteUser(name string) error {
	_, err := collection.DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		return err
	}
	return nil
}
