package main

import (
	"log"
	"net/http"
)

func main() {
	router := SetupBasicRouter()
	loggingRouter := SetupLoggingMiddleWare(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: loggingRouter,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
