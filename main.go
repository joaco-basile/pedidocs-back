package main

import (
	"net/http"

	"github.com/joaco-basile/pedidocs-back/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	routes.SetRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "El server responde")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
