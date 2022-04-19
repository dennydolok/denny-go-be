package domains

import (
	"arc/models"
)

type UserRepositoryDomain interface {
	CreateUsers(user models.User) error
	GetUsers() []models.User
	Verify(email string) (user models.User, err error)
}

type UserService interface {
	CreateUserService(user models.User) error
	GetUsersService() []models.User
	LoginUser(email, password string) (string, int)
}
