package web

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestLogAccess_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Cfg: &Config{
			Log: LogConfig{
				AccessLogs: true,
			},
		},
	}

	ctx := &beecontext.Context{
		Request: &http.Request{
			Method:     "GET",
			RequestURI: "/test",
			Proto:      "HTTP/1.1",
			Host:       "localhost",
		},
	}

	startTime := time.Now()
	statusCode := 200

	app.LogAccess(ctx, &startTime, statusCode)

	// Add assertions here to verify the expected behavior
}
