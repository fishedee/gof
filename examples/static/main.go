package main

import (
	"github.com/fishedee/gof"
	"net/http"
)

func main() {
	router := gof.NewDefaultRouterFactory()
	router.Static("/assets", "../../")

	// Listen and serve on 0.0.0.0:8080
	http.Handle("/", router.Create())
	http.ListenAndServe(":8080", nil)
}
