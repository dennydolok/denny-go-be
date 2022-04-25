package rest

import (
	"arc/config"
	"arc/databases"
	"arc/repositories"
	"arc/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUsersAPI(e *echo.Echo, conf config.Config) {
	db := databases.InitMysql(conf)
	repository := repositories.NewUserRepository(db)
	service := services.NewServiceUsers(conf, repository)
	controller := controllerUser{
		service: service,
	}
	userAPI := e.Group("/users", middleware.Logger())
	userAPI.POST("/login", controller.LoginUserController)
	userAPI.GET("", controller.GetUsersController, middleware.JWT([]byte(conf.SECRET_KEY)))
	userAPI.POST("", controller.CreateUserController)
}
