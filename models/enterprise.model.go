package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Enterprise struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `json:"name"`
	Products []Product            `json:"products"`
	Orders   []Order              `json:"orders"`
	Users    []primitive.ObjectID `json:"users"`
}
