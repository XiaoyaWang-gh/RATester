package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestAccessLog_1(t *testing.T) {
	type args struct {
		r      *AccessLogRecord
		format string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				r: &AccessLogRecord{
					RemoteAddr:     "127.0.0.1",
					RequestTime:    time.Now(),
					RequestMethod:  "GET",
					Request:        "/",
					ServerProtocol: "HTTP/1.1",
					Host:           "localhost",
					Status:         200,
					BodyBytesSent:  1024,
					ElapsedTime:    100 * time.Millisecond,
					HTTPReferrer:   "http://localhost/",
					HTTPUserAgent:  "Go 1.16",
				},
				format: "{{.RemoteAddr}} {{.RequestTime}} {{.RequestMethod}} {{.Request}} {{.ServerProtocol}} {{.Host}} {{.Status}} {{.BodyBytesSent}} {{.ElapsedTime}} {{.HTTPReferrer}} {{.HTTPUserAgent}}",
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			AccessLog(tt.args.r, tt.args.format)
		})
	}
}
