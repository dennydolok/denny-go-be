package services

import (
	"arc/config"
	"arc/domains"
	"arc/helpers"
	"arc/models"
	"net/http"
)

type serviceUser struct {
	c    config.Config
	repo domains.UserRepositoryDomain
}

func (s *serviceUser) CreateUserService(user models.User) error {
	return s.repo.CreateUsers(user)
}

func (s *serviceUser) GetUsersService() []models.User {
	return s.repo.GetUsers()
}

func (s *serviceUser) LoginUser(email, password string) (string, int) {
	user, _ := s.repo.Verify(email)

	if (user.Password != password) || (user == models.User{}) {
		return "", http.StatusUnauthorized
	}

	token, err := helpers.GenerateToken(user.Email, s.c.SECRET_KEY)

	if err != nil {
		return "", http.StatusInternalServerError
	}
	return token, http.StatusOK
}

func NewServiceUsers(c config.Config, repo domains.UserRepositoryDomain) domains.UserService {
	return &serviceUser{
		c:    c,
		repo: repo,
	}
}
