package gof

import (
	"net/http"
	"testing"
)

func TestRecovery(t *testing.T) {
	routerFactory := NewRouterFactory()
	routerFactory.Use(Recovery())
	routerFactory.GET("/", func(w http.ResponseWriter, r *http.Request) {
		panic("123")
	})
	r, _ := http.NewRequest("GET", "/", nil)
	w := &fakeWriter{}
	router := routerFactory.Create()

	router.ServeHTTP(w, r)
	AssertEqual(t, w.Read(), "server internal error")
}
