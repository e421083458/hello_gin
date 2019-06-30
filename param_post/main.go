package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()

	router.Any("/user/*action", func(c *gin.Context) {
		firstName:=c.DefaultPostForm("first_name","wang")
		lastName:=c.PostForm("last_name")
		c.String(http.StatusOK, "%s,%s",firstName,lastName)
	})

	router.Run(":8080")
}
