package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lish.org/ginBlogs/model"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "dtdba"
	username := "root"
	password := "123123"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
