package accesslog

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestFormat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &CommonLogFormatter{}
	entry := &logrus.Entry{}
	entry.Data = logrus.Fields{
		ClientHost:            "127.0.0.1",
		ClientUsername:        "user",
		RequestMethod:         "GET",
		RequestPath:           "/path",
		RequestProtocol:       "HTTP/1.1",
		DownstreamStatus:      200,
		DownstreamContentSize: 100,
		RequestCount:          1,
		RouterName:            "router",
		ServiceURL:            "service",
	}
	entry.Time = time.Now()
	entry.Level = logrus.InfoLevel
	entry.Message = "message"
	entry.Caller = &runtime.Frame{}
	entry.Context = context.Background()

	b, err := f.Format(entry)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "127.0.0.1 - user [01/Jan/2006:15:04:05 -0700] \"GET /path HTTP/1.1\" 200 100 1 \"router\" \"service\" 1ms\n"
	if string(b) != expected {
		t.Errorf("Expected: %s, got: %s", expected, string(b))
	}
}
