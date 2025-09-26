package middleware

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/zeebo/assert"
)

func TestLoggerWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	config := LoggerConfig{
		Format:           "${time_unix} ${time_unix_milli} ${time_unix_micro} ${time_unix_nano} ${time_rfc3339} ${time_rfc3339_nano} ${time_custom} ${id} ${remote_ip} ${uri} ${host} ${method} ${path} ${route} ${protocol} ${referer} ${user_agent} ${status} ${error} ${latency} ${latency_human} ${bytes_in} ${bytes_out} ${header:X-Request-Id} ${query:q} ${form:q} ${custom}",
		CustomTimeFormat: "2006-01-02 15:04:05.000",
		CustomTagFunc: func(c echo.Context, buf *bytes.Buffer) (int, error) {
			buf.WriteString("custom")
			return 0, nil
		},
	}
	// when
	middleware := LoggerWithConfig(config)
	// then
	assert.NotNil(t, middleware)
}
