package routes

import (
	"myGin/controller"

	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.GET("/check", controller.UserCheck) //查看
		user.GET("/list", controller.UserList)   //列表
		user.POST("/add", controller.UserAdd)    //添加
		user.POST("/edit", controller.UserEdit)  //修改
		user.GET("/delete", controller.UserDel)  //删除
	}

	// user := r.Group("/user")
	// {
	// 	user.GET("/check", controller.UserCheck) //查看
	// 	user.GET("/list", controller.UserList)   //列表
	// 	user.POST("/add", controller.UserAdd)    //添加
	// 	user.POST("/edit", controller.UserEdit)  //修改
	// 	user.GET("/delete", controller.UserDel)  //删除
	// }
}
