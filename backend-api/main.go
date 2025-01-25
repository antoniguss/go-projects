package main

import "net/http"

func main() {

	// router := SetupGinRouter()
	//
	// router.Run(":3000")

	mux := SetupBasicRouter()
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	server.ListenAndServe()
}
