package main

import (
	"arc/config"
	"arc/handlers/rest"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.InitConfig()
	e := echo.New()

	rest.RegisterUsersAPI(e, config)
	e.Start(":8080")
}
