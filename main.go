package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func PostData(context *gin.Context) {
	body := context.Request.Body
	context.IndentedJSON(http.StatusOK, body)

	content, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("error reading body")
	}

	context.JSON(200, gin.H{
		"bodyData": string(content),
	})
}

func GetQuery(context *gin.Context) {
	name := context.Query("name")
	age := context.Query("age")

	context.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {

	// Default logger and middleware
	router := gin.Default()

	router.GET("/getdata", GetData)
	router.POST("/postdata", PostData)
	router.GET("/getquerystring", GetQuery)

	router.Run()
}
