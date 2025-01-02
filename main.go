package main

import "github.com/gin-gonic/gin"

func main() {

	// Default logger and middleware
	router := gin.Default()

	router.GET("/getdata", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.Run()
}
