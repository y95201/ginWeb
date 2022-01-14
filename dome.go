package main

import (
	"fmt"
	"os"
)

func main() {
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello word")
	// })
	// //监听端口默认为8080
	// r.Run(":8080")
	path, _ := os.Getwd()
	// if err != nil {
	// 	log.Printf(err)
	// }
	fmt.Println(path)
}
