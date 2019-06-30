package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, fmt.Sprint(err))
	}
	c.String(200, fmt.Sprintf("%#v", person))
}
