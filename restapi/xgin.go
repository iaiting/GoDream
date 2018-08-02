package restapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	fmt.Println("Start gin")
	router := gin.Default()

	router.GET("/v1/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8082")
}
