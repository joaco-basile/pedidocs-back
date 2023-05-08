package routes

import (
	user_service "github.com/joaco-basile/pedidocs-back/service/user.service"
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {
	public := e.Group("")

	public.POST("/user", user_service.Register)
	public.GET("/user", user_service.Login)
	public.PATCH("/user", user_service.Update)
	public.PUT("/user", user_service.Delete)
}
