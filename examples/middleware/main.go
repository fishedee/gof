package main

import (
	"github.com/fishedee/gof"
	"net/http"
)

func main() {
	// Creates a router without any middleware by default
	router := gof.NewRouterFactory()

	// Global logger middleware
	router.Use(gof.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gof.Recovery())

	// Listen and serve on 0.0.0.0:8080
	http.Handle("/", router.Create())
	http.ListenAndServe(":8080", nil)
}
