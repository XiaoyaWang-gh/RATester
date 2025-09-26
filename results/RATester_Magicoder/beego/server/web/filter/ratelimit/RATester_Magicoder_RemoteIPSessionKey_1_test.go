package ratelimit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestRemoteIPSessionKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &context.Context{
		Request: &http.Request{
			Header: http.Header{
				"X-Real-Ip": []string{"192.168.1.1"},
			},
		},
	}
	ip := RemoteIPSessionKey(ctx)
	if ip != "192.168.1.1" {
		t.Errorf("Expected 192.168.1.1, got %s", ip)
	}

	ctx.Request.Header.Set("X-Forwarded-For", "192.168.1.2")
	ip = RemoteIPSessionKey(ctx)
	if ip != "192.168.1.2" {
		t.Errorf("Expected 192.168.1.2, got %s", ip)
	}

	ctx.Request.Header.Del("X-Forwarded-For")
	ctx.Request.RemoteAddr = "192.168.1.3"
	ip = RemoteIPSessionKey(ctx)
	if ip != "192.168.1.3" {
		t.Errorf("Expected 192.168.1.3, got %s", ip)
	}
}
