package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupGinRouter() *gin.Engine {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name", func(ctx *gin.Context) {
		user := ctx.Params.ByName("name")
		message := "Hello, " + user
		ctx.JSON(http.StatusOK, gin.H{"message": message})

	})

	router.POST("/add", func(ctx *gin.Context) {
		var object InputStruct
		if err := ctx.ShouldBind(&object); err != nil {
			logger.Error(fmt.Sprintf("Error binding JSON: %v", err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		result := ResultStruct{Result: *object.Number1 + *object.Number2}
		ctx.JSON(http.StatusOK, result)
	})

	router.POST("/subtract", func(ctx *gin.Context) {
		var object InputStruct
		if err := ctx.ShouldBind(&object); err != nil {
			logger.Error(fmt.Sprintf("Error binding JSON: %v", err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		result := ResultStruct{Result: *object.Number1 - *object.Number2}

		ctx.JSON(http.StatusOK, result)
	})

	router.POST("/multiply", func(ctx *gin.Context) {
		var object InputStruct
		if err := ctx.ShouldBind(&object); err != nil {
			logger.Error(fmt.Sprintf("Error binding JSON: %v", err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		result := ResultStruct{Result: *object.Number1 * *object.Number2}
		ctx.JSON(http.StatusOK, result)
	})

	router.POST("/divide", func(ctx *gin.Context) {
		var object InputStruct
		if err := ctx.ShouldBind(&object); err != nil {
			logger.Error(fmt.Sprintf("Error binding JSON: %v", err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if *object.Number2 == 0 {
			logger.Error("Error: division by zero")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Division by zero"})
			return
		}

		result := ResultStruct{Result: *object.Number1 / *object.Number2}
		ctx.JSON(http.StatusOK, result)
	})

	return router
}
