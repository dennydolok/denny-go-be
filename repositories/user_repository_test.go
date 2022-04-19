package repositories

import (
	"arc/models"
	"regexp"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var dbmock, mock, _ = sqlmock.New()
var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
	Conn:                      dbmock,
	SkipInitializeWithVersion: true,
},
})
var repo = NewUserRepository(db)

func TestCreateUsers(t *testing.T) {
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"email", "password"}).
			AddRow("denny@gmail.com", "1234"))
	res := repo.GetUsers()
	assert.Len(t, res, 1)
}

func TestCreateUser(t *testing.T) {
	repo := NewUserRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`email`,`password`) VALUES (?,?)")).
		WithArgs("test@gmail.com", "1234").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := repo.CreateUsers(models.User{
		Email:    "test@gmail.com",
		Password: "1234",
	})
	assert.NoError(t, err)
}

func TestVerify(t *testing.T) {
	repo := NewUserRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ?")).WillReturnRows(
		sqlmock.NewRows([]string{"email", "password"}).
			AddRow("denny@gmail.com", "1234"),
	)
	_, err := repo.Verify("denny@gmail.com")
	assert.NoError(t, err)
}
