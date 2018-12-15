package main

import (
	"github.com/fishedee/gof"
	"net/http"
)

func loginEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("loginEndpoint"))
}

func submitEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("submitEndpoint"))
}

func readEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("readEndpoint"))
}

func main() {
	router := gof.NewDefaultRouterFactory()

	// Simple group: v1
	router.Group("/v1", func(v1 *gof.RouterFactory) {
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	})

	// Simple group: v2
	router.Group("/v2", func(v2 *gof.RouterFactory) {
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	})

	// Listen and serve on 0.0.0.0:8080
	http.Handle("/", router.Create())
	http.ListenAndServe(":8080", nil)
}
