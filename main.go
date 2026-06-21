package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default() //r is a server used create a router

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"server running",
		})
	})
	r.Run(":8080")
}