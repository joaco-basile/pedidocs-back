package user_service

import (
	"net/http"
	"time"

	m "github.com/joaco-basile/pedidocs-back/models"
	user_repository "github.com/joaco-basile/pedidocs-back/repository/user.repositoy"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserApi struct{}

var id = primitive.NewObjectID()

var TestUser = m.User{
	ID:        id,
	Name:      "jesus",
	Email:     "jesus@gmail.com",
	Password:  "1234",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func Register(c echo.Context) error {

	err := user_repository.RegisterUser(TestUser)

	if err != nil {
		return c.String(http.StatusBadRequest, "No se pudo registrar el usuario")
	}

	return c.String(http.StatusOK, "Se creo el nuevo usuario")
}

func Login(c echo.Context) error {

	_, err := user_repository.LoginUser(TestUser.Name, TestUser.Email, TestUser.Password)

	if err != nil {
		return c.String(http.StatusBadRequest, "fallo el login")
	}
	return c.String(http.StatusOK, "se logueo con exito")
}

func Update(c echo.Context) error {

	var newUser = m.User{
		Email:    "jose@gmail.com",
		Password: "josesito",
	}

	userID := id.String()

	err := user_repository.UpdateUser(newUser, userID)

	if err != nil {
		return c.String(http.StatusBadRequest, "no se pudo actualizar el usuario")
	}

	return c.String(http.StatusOK, "Se actualizo el usuario")
}

func Delete(c echo.Context) error {

	userID := id.String()

	err := user_repository.DseleteUser(userID)

	if err != nil {
		return c.String(http.StatusBadRequest, "no se pudo eliminar el usuario")
	}

	return c.String(http.StatusOK, "se elimino con exito")
}
