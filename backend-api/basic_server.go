package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SetupBasicRouter() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	})

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		pongMsg, _ := json.Marshal(map[string]string{"message": "pong"})
		fmt.Fprintf(w, "%s", pongMsg)
	})

	mux.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		input := &InputStruct{}
		err := json.Unmarshal(data, input)
		if err != nil || input.Number1 == nil || input.Number2 == nil {
			errorMsg, _ := json.Marshal(map[string]string{"error": "invalid input"})
			fmt.Fprintf(w, "%s", errorMsg)
			return
			// w.WriteHeader(http.StatusBadRequest)
		}
		result := ResultStruct{Result: *input.Number1 + *input.Number2}
		resultMsg, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", resultMsg)

		// ctx.JSON(http.StatusOK, result)

		// log.Printf("r.Form: %v\n", r.Form)
		// log.Printf("r.ContentLength: %v\n", r.ContentLength)
		// log.Printf("data: %v\n", string(data))
		// log.Printf("input: %+v\n", input)
		// fmt.Println("---")
	})

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	//
	// router.GET("/user/:name", func(ctx *gin.Context) {
	// 	user := ctx.Params.ByName("name")
	// 	message := "Hello, " + user
	// 	ctx.JSON(http.StatusOK, gin.H{"message": message})
	//
	// })
	//
	// router.POST("/add", func(ctx *gin.Context) {
	// 	var object InputStruct
	// 	if err := ctx.ShouldBind(&object); err != nil {
	// 		log.Printf("Error binding JSON: %v", err)
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 		return
	// 	}
	//
	// 	result := ResultStruct{Result: *object.Number1 + *object.Number2}
	// 	ctx.JSON(http.StatusOK, result)
	// })
	//
	// router.POST("/subtract", func(ctx *gin.Context) {
	// 	var object InputStruct
	// 	if err := ctx.ShouldBind(&object); err != nil {
	// 		log.Printf("Error binding JSON: %v", err)
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 		return
	// 	}
	//
	// 	result := ResultStruct{Result: *object.Number1 - *object.Number2}
	//
	// 	ctx.JSON(http.StatusOK, result)
	// })
	//
	// router.POST("/multiply", func(ctx *gin.Context) {
	// 	var object InputStruct
	// 	if err := ctx.ShouldBind(&object); err != nil {
	// 		log.Printf("Error binding JSON: %v", err)
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 		return
	// 	}
	//
	// 	result := ResultStruct{Result: *object.Number1 * *object.Number2}
	// 	ctx.JSON(http.StatusOK, result)
	// })
	//
	// router.POST("/divide", func(ctx *gin.Context) {
	// 	var object InputStruct
	// 	if err := ctx.ShouldBind(&object); err != nil {
	// 		log.Printf("Error binding JSON: %v", err)
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 		return
	// 	}
	//
	// 	if *object.Number2 == 0 {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Division by zero"})
	// 		return
	// 	}
	//
	// 	result := ResultStruct{Result: *object.Number1 / *object.Number2}
	// 	ctx.JSON(http.StatusOK, result)
	// })

	return mux

}
