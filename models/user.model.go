package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	VIP          bool               `json:"vip"`
	EnterpriseID primitive.ObjectID `bson:"enterpriseID" json:"enterpriseID"`
	CreatedAt    time.Time          `bson:"created_At" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
