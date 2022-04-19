package repositories

import (
	"arc/domains"
	"arc/models"
	"fmt"

	"gorm.io/gorm"
)

type repositoryMysql struct {
	DB *gorm.DB
}

func (r *repositoryMysql) CreateUsers(user models.User) error {
	error := r.DB.Create(&user).Error
	if error != nil {
		return error
	}
	return nil
}

func (r *repositoryMysql) GetUsers() []models.User {
	users := []models.User{}
	r.DB.Find(&users)

	return users
}

func (r *repositoryMysql) Verify(email string) (user models.User, err error) {
	result := r.DB.Where("email = ?", email).Find(&user)
	if result.RowsAffected < 1 {
		err = fmt.Errorf("Email not found")
	}
	return
}

func NewUserRepository(db *gorm.DB) domains.UserRepositoryDomain {
	return &repositoryMysql{
		DB: db,
	}
}
