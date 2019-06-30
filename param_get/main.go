package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()

	router.GET("/user/*action", func(c *gin.Context) {
		firstName:=c.DefaultQuery("first_name","wang")
		lastName:=c.DefaultQuery("last_name","kai")
		c.String(http.StatusOK, "%s,%s",firstName,lastName)
	})

	router.Run(":8080")
}
