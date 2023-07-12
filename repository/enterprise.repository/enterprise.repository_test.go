package enterprise_repository_test

import (
	"testing"
	"time"

	m "github.com/joaco-basile/pedidocs-back/models"
	er "github.com/joaco-basile/pedidocs-back/repository/enterprise.repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Id = primitive.NewObjectID()

func TestCreate(t *testing.T) {
	userId, err := primitive.ObjectIDFromHex("6466367bd6907fd2515c0758")
	if err != nil {
		t.Error(err)
	}
	var EPrueba = m.Enterprise{
		ID:   Id,
		Name: "adidas",
		Products: []m.Product{
			{
				ID:       primitive.NewObjectID(),
				Name:     "buzo con tres lineas",
				Price:    10000,
				Quantity: 2,
			},
			{
				ID:       primitive.NewObjectID(),
				Name:     "remera lisa",
				Price:    2220,
				Quantity: 3,
			}, {
				ID:       primitive.NewObjectID(),
				Name:     "camperon seleccion de italia",
				Price:    80000,
				Quantity: 1,
			},
		},
		Users:  []primitive.ObjectID{userId, primitive.NewObjectID()},
		Orders: []m.Order{},
	}

	err = er.CreateEnterprise(EPrueba)

	if err != nil {
		t.Error("La prueba fallo por que:", err)
	}
	t.Log("¡¡¡La empresa fue creada con exito!!!")
}
func TestRead(t *testing.T) {
	result, err := er.GetEnterprise(Id)
	if err != nil {
		t.Error(err)
	}
	t.Log("Se logro leer la empresa con exito:", result)
}

func TestAddOrder(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("64a6b664ed349be762b4cf7a")
	if err != nil {
		t.Error(err)
	}

	var or = m.Order{
		ID: primitive.NewObjectID(),
		Products: []m.Product{
			{
				ID:       primitive.NewObjectID(),
				Name:     "zapas",
				Price:    19,
				Quantity: 2,
			},
			{
				ID:       primitive.NewObjectID(),
				Name:     "zapas",
				Price:    19,
				Quantity: 2,
			},
		},
		Date:      time.Now(),
		Buyer:     m.Buyer{Name: "juan", Contact: "juan@gmail.com", Location: "las cañas12"},
		OwnerName: "joaco",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = er.AddOrder(or, id)
	if err != nil {
		t.Error(err)
	}

	t.Log("Se añadio el pedido de forma correcta")
}
