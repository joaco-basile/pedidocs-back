package user_repository

import (
	"context"
	"log"

	"github.com/joaco-basile/pedidocs-back/database"
	m "github.com/joaco-basile/pedidocs-back/models"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func loginUser(name string, email string, password string) (m.User, error) {
	var user m.User

	if name != "" {
		err := collection.FindOne(ctx, bson.D{{"name", name}}).Decode(&user)
		if err != nil {
			log.Panic(err)
		}
	}

	if email != "" {
		err := collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
		if err != nil {
			log.Panic(err)
		}
	}

	isDisMatch := user.Password != password

	if isDisMatch {
		return _, nil
	}

	return user, nil

}

func registerUser(user m.User) error {

}
