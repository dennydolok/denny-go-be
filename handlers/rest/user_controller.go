package rest

import (
	"arc/domains"
	"arc/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controllerUser struct {
	service domains.UserService
}

func (re *controllerUser) CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := re.service.CreateUserService(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":     http.StatusInternalServerError,
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":     http.StatusCreated,
		"messages": "success",
		"users":    user,
	})
}

func (re *controllerUser) GetUsersController(c echo.Context) error {
	users := re.service.GetUsersService()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     http.StatusOK,
		"messages": "success",
		"users":    users,
	})
}

func (re *controllerUser) LoginUserController(c echo.Context) error {
	user := make(map[string]interface{})
	c.Bind(&user)
	token, status := re.service.LoginUser(user["email"].(string), user["password"].(string))
	if status != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "Unable to login",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token":    token,
	})
}
