package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context)  {
	fmt.Println("getting")
}

func posting(c *gin.Context)  {
	fmt.Println("posting")
}

func putting(c *gin.Context)  {
	fmt.Println("putting")
}

func deleting(c *gin.Context)  {
	fmt.Println("deleting")
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

