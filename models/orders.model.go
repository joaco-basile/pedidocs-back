package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Product   string             `json:"product"`
	Quantity  int                `json:"quantity"`
	BuyerName string             `json:"buyer_name"`
	Contact   string             `json:"contact"`
	Date      time.Time          `json:"date"`

	CreatedAt time.Time `bson:"created_At" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
