package databases

import (
	"arc/config"
	"arc/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(conf config.Config) *gorm.DB {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	DB, err := gorm.Open(mysql.Open(connection))
	if err != nil {
		fmt.Println("Cannot connect to database : ", err)
	}
	DB.AutoMigrate(&models.User{})
	return DB
}
