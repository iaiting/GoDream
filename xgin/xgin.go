package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	fmt.Println("this is GetHandler")
}

func PutHandler(c *gin.Context) {
	fmt.Println("this is PutHandler")
}

func PostHandler(c *gin.Context) {
	fmt.Println("this is PostHandler")
}

func DeleteHandler(c *gin.Context) {
	fmt.Println("this is DeleteHandler")
}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/get", GetHandler)
	router.PUT("/put", PutHandler)
	router.POST("/post", PostHandler)
	router.DELETE("/delete", DeleteHandler)
	router.Run()
}
