package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Buyer struct {
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Location string `json:"location"`
}

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Products []Product          `json:"products"`
	Date     time.Time          `json:"date"`
	Buyer    Buyer              `json:"buyer"`

	OwnerName string    `json:"owner_name"`
	CreatedAt time.Time `bson:"created_At" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
