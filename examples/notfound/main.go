package main

import (
	"github.com/fishedee/gof"
	"net/http"
)

func main() {
	router := gof.NewDefaultRouterFactory()
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("404 not found by hello world"))
	})

	// Listen and serve on 0.0.0.0:8080
	http.Handle("/", router.Create())
	http.ListenAndServe(":8080", nil)
}
