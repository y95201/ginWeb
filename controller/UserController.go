package controller

//controller/Index.go

import (
	"myGin/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UserCheck(context *gin.Context) {

	users := models.GetProjects(context)
	context.JSON(200, users)
	// router := gin.Default()
	// users := "123"
	// // 加载静态文件
	// router.Static("/public", "./public")
	// router.LoadHTMLGlob("views/*")
	// // 加载静态文件
	// context.HTML(http.StatusOK, "/index.html", gin.H{"title": users})
}
func UserList(context *gin.Context) {

	users := models.GetList(context)
	context.JSON(200, users)

}
func UserAdd(context *gin.Context) {

	users := models.GetAdds(context)
	context.JSON(200, users)

}
func UserEdit(context *gin.Context) {

	users := models.GetUpdate(context)
	context.JSON(200, users)

}
func UserDel(context *gin.Context) {

	users := models.GetDel(context)
	context.JSON(200, users)

}
