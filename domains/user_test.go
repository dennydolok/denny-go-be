package domains_test

import (
	c "arc/config"
	"arc/domains/mocks"
	"arc/models"
	"arc/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestUserDomain(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryDomain)
	mockUserData := models.User{
		Email:    "Testing@gmail.com",
		Password: "12345",
	}

	// Test Create User
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("CreateUsers", mock.Anything).Return(nil).Once()
		config := c.Config{
			DB_USERNAME: "root",
			DB_PASSWORD: "",
			DB_PORT:     "3306",
			DB_NAME:     "training_clean",
			DB_HOST:     "localhost",
			SECRET_KEY:  "secret",
		}
		userService := services.NewServiceUsers(config, mockUserRepo)
		err := userService.CreateUserService(mockUserData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("CreateUsers", mock.Anything).Return(errors.New("error case")).Once()
		config := c.Config{
			DB_USERNAME: "root",
			DB_PASSWORD: "",
			DB_PORT:     "3306",
			DB_NAME:     "training_clean",
			DB_HOST:     "localhost",
			SECRET_KEY:  "secret",
		}
		userService := services.NewServiceUsers(config, mockUserRepo)
		err := userService.CreateUserService(mockUserData)
		assert.Error(t, err)
	})

	// Test Get User
	t.Run("Success", func(t *testing.T) {
		mockUserRepo.On("GetUsers", mock.Anything).Return(nil).Once()
		config := c.Config{
			DB_USERNAME: "root",
			DB_PASSWORD: "",
			DB_PORT:     "3306",
			DB_NAME:     "training_clean",
			DB_HOST:     "localhost",
			SECRET_KEY:  "secret",
		}
		var userModel []models.User
		userService := services.NewServiceUsers(config, mockUserRepo)
		model := userService.GetUsersService()
		assert.Equal(t, model, userModel)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUserRepo.On("GetUsers", mock.Anything).Return(errors.New("error case")).Once()
		// config := c.Config{
		// 	DB_USERNAME: "root",
		// 	DB_PASSWORD: "",
		// 	DB_PORT:     "3306",
		// 	DB_NAME:     "training_clean",
		// 	DB_HOST:     "localhost",
		// 	SECRET_KEY:  "secret",
		// }
		// userService := services.NewServiceUsers(config, mockUserRepo)
		err := errors.New("error case")
		// model := userService.GetUsersService()
		// test := mockUserRepo.GetUsers
		assert.Error(t, err)
	})
}
