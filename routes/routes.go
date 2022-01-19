package routes

import (
	"myGin/controller"

	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	// gin.DisableConsoleColor()
	// // 记录到文件。
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// // 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

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
