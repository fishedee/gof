package gof

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	green        = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white        = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow       = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red          = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue         = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta      = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan         = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

type LoggerConfig struct {
	Output       io.Writer
	DisbaleColor bool
}

type logFormatterParams struct {
	Request      *http.Request
	TimeStamp    time.Time
	Latency      time.Duration
	ClientIP     string
	Method       string
	Path         string
	DisbaleColor bool
}

func DisableConsoleColor() {
	disableColor = true
}

func Logger() RouterMiddleware {
	return LoggerWithConfig(LoggerConfig{
		Output:       os.Stdout,
		DisbaleColor: disableColor,
	})
}

func LoggerWithConfig(conf LoggerConfig) RouterMiddleware {

	disableColor := false
	if _, ok := conf.Output.(*os.File); !ok ||
		os.Getenv("TERM") == "dumb" ||
		conf.DisbaleColor {
		disableColor = true
	}

	return func(prev RouterMiddlewareContext) RouterMiddlewareContext {
		last := prev.Handler.(func(w http.ResponseWriter, r *http.Request, p RouterParam))
		return RouterMiddlewareContext{
			Data: prev.Data,
			Handler: func(w http.ResponseWriter, r *http.Request, p RouterParam) {
				start := time.Now()
				path := r.URL.Path
				raw := r.URL.RawQuery
				last(w, r, p)
				param := logFormatterParams{
					Request:      r,
					DisbaleColor: disableColor,
				}

				// Stop timer
				param.TimeStamp = time.Now()
				param.Latency = param.TimeStamp.Sub(start)

				param.ClientIP = getClientIP(r)
				param.Method = r.Method

				if raw != "" {
					path = path + "?" + raw
				}

				param.Path = path

				fmt.Fprintf(conf.Output, logFormatter(param))
			},
		}
	}
}

func getClientIP(r *http.Request) string {
	ips := getClientProxy(r)
	if len(ips) > 0 && ips[0] != "" {
		rip, _, err := net.SplitHostPort(ips[0])
		if err != nil {
			rip = ips[0]
		}
		return rip
	}
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}
	return r.RemoteAddr
}

func getClientProxy(r *http.Request) []string {
	if ips := r.Header.Get("X-Forwarded-For"); ips != "" {
		return strings.Split(ips, ",")
	}
	return []string{}
}

func logFormatter(param logFormatterParams) string {
	var methodColor string
	if param.DisbaleColor == false {
		methodColor = colorForMethod(param.Method)
	}

	return fmt.Sprintf("[GOF] %13v | %15s |%s %-7s %s %s|%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, reset,
		param.Path,
	)
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}
