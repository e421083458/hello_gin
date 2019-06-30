package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("example","123456")
		ipList:=[]string{
			"127.0.0.1",
		}
		isMatched:=false
		for _, host := range ipList {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched{
			c.String(200,fmt.Sprintf("%v, not in iplist", c.ClientIP()))
			c.Abort()
		}
	}
}

func main() {
	r := gin.New()
	r.Use(IPAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	r.Run(":8080")
}