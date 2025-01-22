package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/user/:name", func(ctx *gin.Context) {
		user := ctx.Params.ByName("name")
		message := "Hello, " + user
		ctx.JSON(http.StatusOK, gin.H{"message": message})

	})

	router.Run(":3000")
}
