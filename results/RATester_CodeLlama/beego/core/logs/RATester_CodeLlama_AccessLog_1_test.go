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
		BodyBytesSent:  100,
		ElapsedTime:    100 * time.Millisecond,
		HTTPReferrer:   "http://localhost/",
		HTTPUserAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
		RemoteUser:     "test",
	}
	format := "%{remote_addr}i %{request_time}t %{request_method}t %{request}t %{server_protocol}t %{host}t %{status}t %{body_bytes_sent}t %{elapsed_time}t %{http_referrer}t %{http_user_agent}t %{remote_user}t"
	AccessLog(r, format)
}
