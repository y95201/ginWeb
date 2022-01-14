package main

import (
	"myGin/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	router := gin.Default()

	routes.Load(router)

	router.Run(":8080")

}
