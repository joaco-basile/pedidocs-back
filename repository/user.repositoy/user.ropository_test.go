package user_repository_test

import (
	"fmt"
	"testing"
	"time"

	m "github.com/joaco-basile/pedidocs-back/models"
	user_repo "github.com/joaco-basile/pedidocs-back/repository/user.repositoy"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId = primitive.NewObjectID()

func TestRegister(t *testing.T) {
	user := m.User{
		ID:        userId,
		Name:      "Jesus",
		Email:     "jesus.matiz@micorreo.com",
		Password:  "123",
		VIP:       true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := user_repo.RegisterUser(user)

	if err != nil {
		t.Error("El usuario no se pudo registrar")
		t.Fail()
	} else {
		t.Log("la prueba de registrarse finalizo con exito")
	}

}

func TestLogin(t *testing.T) {

	user, err := user_repo.LoginUser("Jesus", "123")

	if err != nil {
		t.Error("El login ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba de login finalizo con exito!", user)
	}
	fmt.Println("si llega hasta abajo en el test")
}

func TestUpdate(t *testing.T) {

	user := m.User{
		Name:  "Jesus Matiz",
		Email: "jesus.matiz.prg@gmail.com",
	}

	err := user_repo.UpdateUser(user, userId)

	if err != nil {
		t.Error("Error al tratar de actualizar el usuario")
		t.Fail()
	} else {
		t.Log("La prueba de actualización finalizo con exito!")
	}
}

func TestDelete(t *testing.T) {

	err := user_repo.DeleteUser(userId)

	if err != nil {
		t.Error("Error al tratar de eliminar el usuario")
		t.Fail()
	} else {
		t.Log("La prueba de eliminación finalizo con exito!")
	}
}
