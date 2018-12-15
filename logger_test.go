package gof

import (
	"net/http"
	"testing"
)

func TestLogger(t *testing.T) {
	routerFactory := NewRouterFactory()
	routerFactory.Use(Logger())
	routerFactory.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello Fish"))
	})
	r, _ := http.NewRequest("GET", "/", nil)
	w := &fakeWriter{}
	router := routerFactory.Create()

	router.ServeHTTP(w, r)
	AssertEqual(t, w.Read(), "Hello Fish")
}
