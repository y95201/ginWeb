package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var err error

type User struct {
	Username string
}

func GetProjects(c *gin.Context) (users []User) {
	db := InitDB()
	db.AutoMigrate(&User{})
	db.Where("id = ?", "3").First(&users)
	return users
}

func InitDB() *gorm.DB {

	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/think?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	return db

}
