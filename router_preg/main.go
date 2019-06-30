package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()

	router.GET("/user/*action", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	router.Run(":8080")
}
