package main

import (
	"github.com/fishedee/gof"
	"net/http"
)

func doSomething(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("something have done"))
}

func main() {
	// Disable Console Color
	// gof.DisableConsoleColor()

	// Creates a gof router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gof.NewDefaultRouterFactory()

	router.GET("/someGet", doSomething)
	router.POST("/somePost", doSomething)
	router.PUT("/somePut", doSomething)
	router.DELETE("/someDelete", doSomething)
	router.PATCH("/somePatch", doSomething)
	router.HEAD("/someHead", doSomething)
	router.OPTIONS("/someOptions", doSomething)
	router.Any("/someAny", doSomething)

	// Listen and serve on 0.0.0.0:8080
	http.Handle("/", router.Create())
	http.ListenAndServe(":8080", nil)
}
