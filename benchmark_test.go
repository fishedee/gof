package gof

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func benchmarkRouterBasic(b *testing.B, insertData []string, findData []string) {
	routerFactory := NewRouterFactory()
	doNothing := func(w http.ResponseWriter, r *http.Request, param RouterParam) {
	}
	routerFactory.NotFound(doNothing)
	for _, data := range insertData {
		routerFactory.GET(data, doNothing)
	}

	r, _ := http.NewRequest("GET", "", nil)
	w := &fakeWriter{}
	router := routerFactory.Create()

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		single := findData[i%len(findData)]
		r.URL.Path = single
		router.ServeHTTP(w, r)
	}
}

func benchmarkGinBasic(b *testing.B, insertData []string, findData []string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	doNothing := func(c *gin.Context) {
	}
	for _, data := range insertData {
		router.GET(data, doNothing)
	}

	r, _ := http.NewRequest("GET", "", nil)
	w := &fakeWriter{}

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		single := findData[i%len(findData)]
		r.URL.Path = single
		router.ServeHTTP(w, r)
	}
}

func BenchmarkRouterShort(b *testing.B) {
	testUrl := "/abc"
	benchmarkRouterBasic(b, []string{testUrl}, []string{testUrl})
}

func BenchmarkGinShort(b *testing.B) {
	testUrl := "/abc"
	benchmarkGinBasic(b, []string{testUrl}, []string{testUrl})
}

func BenchmarkRouterLong(b *testing.B) {
	testUrl := "/abc/12312313/adf/asdf/asdf/asdf/sdaf/asdf/abc/12312313/adf/asdf/asdf/asdf/sdaf/asdf/"
	benchmarkRouterBasic(b, []string{testUrl}, []string{testUrl})
}

func BenchmarkGinLong(b *testing.B) {
	testUrl := "/abc/12312313/adf/asdf/asdf/asdf/sdaf/asdf/abc/12312313/adf/asdf/asdf/asdf/sdaf/asdf/"
	benchmarkGinBasic(b, []string{testUrl}, []string{testUrl})
}

func BenchmarkRouterParam(b *testing.B) {
	insertUrl := "/user/:userId"
	findUrl := "/user/123"
	benchmarkRouterBasic(b, []string{insertUrl}, []string{findUrl})
}

func BenchmarkGinParam(b *testing.B) {
	insertUrl := "/user/:userId"
	findUrl := "/user/123"
	benchmarkGinBasic(b, []string{insertUrl}, []string{findUrl})
}

func BenchmarkRouterParamLong(b *testing.B) {
	insertUrl := "/user/:userId"
	findUrl := "/user/123/adsfasdfadsfasdfasdfadsf/zcvczxcxzvzvcx"
	benchmarkRouterBasic(b, []string{insertUrl}, []string{findUrl})
}

func BenchmarkGinParamLong(b *testing.B) {
	insertUrl := "/user/:userId"
	findUrl := "/user/123/adsfasdfadsfasdfasdfadsf/zcvczxcxzvzvcx"
	benchmarkGinBasic(b, []string{insertUrl}, []string{findUrl})
}
