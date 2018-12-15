package main

import (
	"fmt"
	"github.com/fishedee/gof"
	"net/http"
	"net/url"
)

func ParseFormAndQueryMiddleware() gof.RouterMiddleware {
	return func(prev gof.RouterMiddlewareContext) gof.RouterMiddlewareContext {
		last, isOk := prev.Handler.(func(form url.Values) string)
		if isOk == false {
			return prev
		}
		return gof.RouterMiddlewareContext{
			Data: prev.Data,
			Handler: func(w http.ResponseWriter, r *http.Request, param gof.RouterParam) {
				err := r.ParseForm()
				if err != nil {
					panic(err)
				}
				form := r.Form
				for _, singleParam := range param {
					form.Add(singleParam.Key, singleParam.Value)
				}
				data := last(form)
				w.WriteHeader(200)
				w.Write([]byte(data))
			},
		}
	}
}

func main() {
	router := gof.NewDefaultRouterFactory()
	router.Use(ParseFormAndQueryMiddleware())

	//standard style get form value and param value
	//try http://localhost:8080/standard/fish?a=3&b=4
	router.GET("/standard/:userId", func(w http.ResponseWriter, r *http.Request, param gof.RouterParam) {
		r.ParseForm()
		fmt.Printf("form:%v,param:%v\n", r.Form, param)
		w.WriteHeader(200)
		w.Write([]byte("success"))
	})

	//user define style get form value and param value
	//try http://localhost:8080/user/fish?a=3&b=4
	router.GET("/user/:userId", func(form url.Values) string {
		fmt.Printf("form and param:%v\n", form)
		return "success2"
	})

	// Listen and serve on 0.0.0.0:8080
	http.Handle("/", router.Create())
	http.ListenAndServe(":8080", nil)
}
