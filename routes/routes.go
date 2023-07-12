package routes

import (
	enterprise_service "github.com/joaco-basile/pedidocs-back/service/enterprise.service"
	user_service "github.com/joaco-basile/pedidocs-back/service/user.service"
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {
	public := e.Group("")

	//rutas para los metodos de usuarios
	public.POST("/user", user_service.Register)
	public.GET("/user", user_service.Login)
	public.PATCH("/user", user_service.Update)
	public.PUT("/user/:id", user_service.Delete)

	//rutas para los metodos de empresas
	public.POST("/enterprise", enterprise_service.CreateEnterprise)

}
