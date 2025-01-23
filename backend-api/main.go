package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InputStruct struct {
	Number1 *int `json:"number1" binding:"required"`
	Number2 *int `json:"number2" binding:"required"`
}

type ResultStruct struct {
	Result int `json:"result"`
}

func setupRouter() *gin.Engine {

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
			log.Printf("Error binding JSON: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		result := ResultStruct{Result: *object.Number1 + *object.Number2}
		ctx.JSON(http.StatusOK, result)
	})

	router.POST("/subtract", func(ctx *gin.Context) {
		var object InputStruct
		if err := ctx.ShouldBind(&object); err != nil {
			log.Printf("Error binding JSON: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		result := ResultStruct{Result: *object.Number1 - *object.Number2}

		ctx.JSON(http.StatusOK, result)
	})

	router.POST("/multiply", func(ctx *gin.Context) {
		var object InputStruct
		if err := ctx.ShouldBind(&object); err != nil {
			log.Printf("Error binding JSON: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		result := ResultStruct{Result: *object.Number1 * *object.Number2}
		ctx.JSON(http.StatusOK, result)
	})

	router.POST("/divide", func(ctx *gin.Context) {
		var object InputStruct
		if err := ctx.ShouldBind(&object); err != nil {
			log.Printf("Error binding JSON: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if *object.Number2 == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Division by zero"})
			return
		}

		result := ResultStruct{Result: *object.Number1 / *object.Number2}
		ctx.JSON(http.StatusOK, result)
	})

	return router
}

func main() {

	router := setupRouter()

	router.Run(":3000")
}
