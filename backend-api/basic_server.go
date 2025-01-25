package main

import (
	"encoding/json"
	"fmt"
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
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
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

	mux.HandleFunc("POST /subtract", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		input := &InputStruct{}
		err := json.Unmarshal(data, input)
		if err != nil || input.Number1 == nil || input.Number2 == nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		result := ResultStruct{Result: *input.Number1 - *input.Number2}
		resultMsg, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", resultMsg)

		// ctx.JSON(http.StatusOK, result)

		// log.Printf("r.Form: %v\n", r.Form)
		// log.Printf("r.ContentLength: %v\n", r.ContentLength)
		// log.Printf("data: %v\n", string(data))
		// log.Printf("input: %+v\n", input)
		// fmt.Println("---")
	})

	mux.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		input := &InputStruct{}
		err := json.Unmarshal(data, input)
		if err != nil || input.Number1 == nil || input.Number2 == nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		result := ResultStruct{Result: *input.Number1 * *input.Number2}
		resultMsg, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", resultMsg)

		// ctx.JSON(http.StatusOK, result)

		// log.Printf("r.Form: %v\n", r.Form)
		// log.Printf("r.ContentLength: %v\n", r.ContentLength)
		// log.Printf("data: %v\n", string(data))
		// log.Printf("input: %+v\n", input)
		// fmt.Println("---")
	})

	mux.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		input := &InputStruct{}
		err := json.Unmarshal(data, input)
		if err != nil || input.Number1 == nil || input.Number2 == nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		if *input.Number2 == 0 {
			http.Error(w, "invalid input (division by zero)", http.StatusBadRequest)
			return
		}

		result := ResultStruct{Result: *input.Number1 / *input.Number2}
		resultMsg, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", resultMsg)

		// ctx.JSON(http.StatusOK, result)

		// log.Printf("r.Form: %v\n", r.Form)
		// log.Printf("r.ContentLength: %v\n", r.ContentLength)
		// log.Printf("data: %v\n", string(data))
		// log.Printf("input: %+v\n", input)
		// fmt.Println("---")
	})

	return mux

}
