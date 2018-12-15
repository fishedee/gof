package main

import (
	"fmt"
	"github.com/fishedee/gof"
	"net/http"
)

func main() {
	router := gof.NewDefaultRouterFactory()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, param gof.RouterParam) {
		name := param[0].Value
		fmt.Fprintf(w, "name: %v", name)
	})

	// However, this one will match /user/john/post and also /user/john/send
	router.GET("/user/:name/:action", func(w http.ResponseWriter, r *http.Request, param gof.RouterParam) {
		name := param[0].Value
		action := param[1].Value
		fmt.Fprintf(w, "%v is %v", name, action)
	})

	// and this one will match /user/fish/post and also /user/fish/send
	router.GET("/user/fish/:action", func(w http.ResponseWriter, r *http.Request, param gof.RouterParam) {
		action := param[0].Value
		fmt.Fprintf(w, "friend fish is %v", action)
	})

	http.Handle("/", router.Create())
	http.ListenAndServe(":8080", nil)
}
