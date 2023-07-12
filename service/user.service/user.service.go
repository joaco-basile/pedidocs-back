package user_service

import (
	"net/http"
	"time"

	m "github.com/joaco-basile/pedidocs-back/models"
	user_repository "github.com/joaco-basile/pedidocs-back/repository/user.repositoy"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c echo.Context) error {
	var json map[string]interface{} = map[string]interface{}{}

	err := c.Bind(&json)
	if err != nil {
		return err
	}

	var user = m.User{
		ID:        primitive.NewObjectID(),
		Name:      json["name"].(string),
		Email:     json["email"].(string),
		Password:  json["password"].(string),
		VIP:       false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = user_repository.RegisterUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {

	var json map[string]interface{} = map[string]interface{}{}

	err := c.Bind(&json)
	if err != nil {
		return err
	}

	user, err := user_repository.LoginUser(json["name"].(string), json["password"].(string))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func Update(c echo.Context) error {
	var json map[string]interface{} = map[string]interface{}{}
	err := c.Bind(&json)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUser := m.User{
		Name:     json["name"].(string),
		Password: json["password"].(string),
		Email:    json["email"].(string),
	}

	id, err := primitive.ObjectIDFromHex(json["id"].(string))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = user_repository.UpdateUser(newUser, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, "se actualizo el usuario con exito")
}

func Delete(c echo.Context) error {

	id := c.Param("id")

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = user_repository.DeleteUser(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, "se elimino el usuario con exito")
}
