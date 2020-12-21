package main

import "github.com/mjrpereira/cnbv/handlers"

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Pong",
	//	})
	//})
	//v1 := r.Group("/api/v1")
	//{
	//	v1.POST("/user", s)
	//}

	//r.Run()
	handlers.Run()
}
