package user_repository

import (
	"context"
	"errors"
	"time"

	"github.com/joaco-basile/pedidocs-back/database"
	m "github.com/joaco-basile/pedidocs-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func CheckExist(email string) bool {
	var existUser m.User
	collection.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&existUser)
	return existUser.Email == email
}

func LoginUser(email string, password string) (m.User, error) {
	var user m.User

	err := collection.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)
	if user.Password != password {
		err = errors.New("la contraseña no coincide")
		return m.User{}, err
	}
	if err != nil {
		return m.User{}, err
	}

	return user, nil

}

func RegisterUser(user m.User) error {
	isExist := CheckExist(user.Email)
	if isExist {
		err := errors.New("ya existe un usuario con ese gmail, pruebe mejor con otro")
		return err
	}

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

func AddToEnterprise(userID primitive.ObjectID, enterpriseID primitive.ObjectID) error {
	var err error

	filter := bson.M{"_id": userID}

	update := bson.M{
		"$set": bson.M{
			"enterpriseID": enterpriseID,
			"updated_at":   time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
