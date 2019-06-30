package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main()  {
	router := gin.Default()

	router.Any("/user/*action", func(c *gin.Context) {
		bodyBytes,err:=ioutil.ReadAll(c.Request.Body)
		if err!=nil{
			c.String(http.StatusBadRequest, err.Error())
		}
		c.String(http.StatusOK, string(bodyBytes))
	})

	router.Run(":8080")
}
