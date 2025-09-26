package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestAccessLog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &AccessLogRecord{
		RemoteAddr:     "127.0.0.1",
		RequestTime:    time.Now(),
		RequestMethod:  "GET",
		Request:        "/",
		ServerProtocol: "HTTP/1.1",
		Host:           "localhost",
		Status:         200,
		BodyBytesSent:  1024,
		ElapsedTime:    time.Second,
		HTTPReferrer:   "http://example.com",
		HTTPUserAgent:  "Mozilla/5.0",
	}
	format := "%{remote_addr} %{request_time} %{request_method} %{request} %{server_protocol} %{host} %{status} %{body_bytes_sent} %{elapsed_time} %{http_referrer} %{http_user_agent}"
	AccessLog(r, format)
}
