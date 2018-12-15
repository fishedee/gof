# Gof Web Framework

Gof is a web framework written in Go (Golang). It features a Martini-like API with much better performance, up to 45 times faster to [martini](https://github.com/go-martini/martini) , and  up to 15% faster to [gin](https://github.com/gin-gonic/gin) If you need performance and good productivity, you will love Gof.

Except to performance , gof alos have powerful scalablity,you can build your own request param with very less cost.

## Feactures

- [x] Zero allocation router.
- [x] Still the fastest http router and framework. Even faster than Gin
- [x] Complete suite of unit tests
- [x] API frozen, new releases will not break your code.
- [x] Powerful scalablity,you can build your own request param with very less cost.
- [x] Accurate Match,gof can excactly decide /a/b and /a/(:userId) router in evey scene.

## Benchmarks

[See all benchmarks](/benchmark_test.go)

Benchmark name                              | (1)        | (2)         |
--------------------------------------------|-----------:|------------:|
BenchmarkGofShort-4                  | 30000000  |  48.7 ns/op |
BenchmarkGinShort-4                  | 20000000  |  68.8 ns/op |
BenchmarkGofLong-4                  | 30000000  |   50.8 ns/op |
BenchmarkGinLong-4                  | 20000000  |  72.3 ns/op |
BenchmarkGofParam-4                  | 20000000  |  96.8 ns/op |
BenchmarkGinParam-4                  | 20000000  |  76.4 ns/op |
BenchmarkGofParamLong-4                  | 20000000  |  105  ns/op |
BenchmarkGinParamLong-4                  | 100000  |  143288  ns/op |

- (1): Total Repetitions achieved in constant time, higher means more confident result
- (2): Single Repetition Duration (ns/op), lower is better

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/fishedee/gof
```

## API Examples

### Using GET, POST, PUT, PATCH, DELETE and OPTIONS

```go
func main() {
    // Disable Console Color
    // gof.DisableConsoleColor()

    // Creates a gof router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gof.NewDefaultRouterFactory()

    router.GET("/someGet", getting)
    router.POST("/somePost", posting)
    router.PUT("/somePut", putting)
    router.DELETE("/someDelete", deleting)
    router.PATCH("/somePatch", patching)
    router.HEAD("/someHead", head)
    router.OPTIONS("/someOptions", options)
    router.ANY("/someAny",any)

    http.Handle('/',router.Create())
}
```

### Parameters in path

```go
func main() {
    router := gof.NewDefaultRouterFactory()

    // This handler will match /user/john but will not match /user/ or /user
    router.GET("/user/:name", func(w http.ResponseWriter,r * http.Request,param gof.RouterParam) {
        name := param[0].Value
        fmt.Fprintf(w,"name: %v",name)
    })

    // However, this one will match /user/john/post and also /user/john/send
    router.GET("/user/:name/:action", func(w http.ResponseWriter,r * http.Request,param gof.RouterParam) {
        name := param[0].Value
        action := param[1].Value
        fmt.Fprintf(w,"%v is  %v",name,action)
    })

    // and this one will match /user/fish/post and also /user/fish/send
    router.GET("/user/fish/:action", func(w http.ResponseWriter,r * http.Request,param gof.RouterParam) {
        action := param[0].Value
        fmt.Fprintf(w,"friend fish is  %v",name,action)
    })

    http.Handle('/',router.Create())
}
```

### Serving static files

```go
func main() {
    router := gof.NewDefaultRouterFactory()
    router.Static("/assets", "./assets")

    // Listen and serve on 0.0.0.0:8080
    http.Handle('/',router.Create())
}
```

### Serving not found

```go
func main() {
    router := gof.NewDefaultRouterFactory()
    router.NotFound(func(w http.ResponseWriter, r *http.Request) {
    	w.Write([]byte("404 not found by hello world"))
    })

    // Listen and serve on 0.0.0.0:8080
    http.Handle('/',router.Create())
}
```

### Grouping routes

```go
func main() {
    router := gof.NewDefaultRouterFactory()

    // Simple group: v1
    router.Group("/v1",func(v1 *gof.RouterFactory) {
    	v1.POST("/login", loginEndpoint)
        v1.POST("/submit", submitEndpoint)
        v1.POST("/read", readEndpoint)
    })

    // Simple group: v2
    router.Group("/v2",func(v2 *gof.RouterFactory) {
    	v2.POST("/login", loginEndpoint)
        v2.POST("/submit", submitEndpoint)
        v2.POST("/read", readEndpoint)
    })

    // Listen and serve on 0.0.0.0:8080
    http.Handle('/',router.Create())
}
```


### Using middleware

```go
func main() {
    // Creates a router without any middleware by default
    r := gof.NewRouterFactory()

    // Global middleware
    // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
    r.Use(gof.Logger())

    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gof.Recovery())

    // Listen and serve on 0.0.0.0:8080
    http.Handle('/',router.Create())
}
```

## License

MIT licensed. See the LICENSE file for details.