package enterprise_service

import (
	"fmt"
	"net/http"

	m "github.com/joaco-basile/pedidocs-back/models"
	enterprise_repository "github.com/joaco-basile/pedidocs-back/repository/enterprise.repository"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEnterprise(c echo.Context) error {
	//Struct de datos que debe recibir
	type data struct {
		Name  string   `json:"name"`
		Users []string `json:"users"`
	}

	//Guardando los datos en d
	d := new(data)
	err := c.Bind(&d)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(d)

	//String To Objet ID
	var UsersIds []primitive.ObjectID
	for _, ids := range d.Users {
		fmt.Println(ids)
		id, err := primitive.ObjectIDFromHex(ids)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		UsersIds = append(UsersIds, id)
	}

	//Creando una enterprise con los datos recibidos
	var enterprise = m.Enterprise{
		ID:    primitive.NewObjectID(),
		Name:  d.Name,
		Users: UsersIds,
	}
	//Se llama a la funcion en repository para que la almacene en la Db
	err = enterprise_repository.CreateEnterprise(enterprise)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error al crear la empresa")
	}
	//Respuesta del server si esta todo Ok
	return c.String(http.StatusOK, "Se creo la empresa")
}

func AddOrder(c echo.Context) error {
	//Se crea un tipo de dato y una instancia del mismo
	type data struct {
		IdEnterprise string
		NewOrder     m.Order
	}
	d := new(data)

	//Se guardan los datos de la peticion
	err := c.Bind(d)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	//Transformando de string a objetId
	id, err := primitive.ObjectIDFromHex(d.IdEnterprise)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	//Se llama al repository
	enterprise_repository.AddOrder(d.NewOrder, id)

	//Se devuelve una respuesta al pedido
	return c.String(http.StatusOK, "¡¡El pedido fue guardado con exito!!")
}

func GetEnterprise(c echo.Context) error {
	id := c.Param("enterpriseID")
	objetID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.Error(err)
	}

	enterpriseData, err := enterprise_repository.GetEnterprise(objetID)
	if err != nil {
		c.Error(err)
	}

	return c.JSON(http.StatusOK, enterpriseData)
}
