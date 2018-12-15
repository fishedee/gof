package gof

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

func getStackTrace() string {
	stack := ""
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
	}
	return stack
}

func Recovery() RouterMiddleware {
	return RecoveryWithWriter(os.Stderr)
}

func RecoveryWithWriter(out io.Writer) RouterMiddleware {
	var logger *log.Logger
	logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
	return func(prev RouterMiddlewareContext) RouterMiddlewareContext {
		last := prev.Handler.(func(w http.ResponseWriter, r *http.Request, param RouterParam))
		return RouterMiddlewareContext{
			Data: prev.Data,
			Handler: func(w http.ResponseWriter, r *http.Request, param RouterParam) {
				defer func() {
					err := recover()
					if err != nil {
						stack := getStackTrace()
						logger.Printf("[Recovery] panic recovered:\n%s\n%s", err, stack)
						w.WriteHeader(500)
						w.Write([]byte("server internal error"))
					}
				}()
				last(w, r, param)
			},
		}
	}
}
