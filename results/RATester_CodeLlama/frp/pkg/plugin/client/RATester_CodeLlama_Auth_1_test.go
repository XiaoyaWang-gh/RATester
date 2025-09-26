package client

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hp := &HTTPProxy{
		opts: &v1.HTTPProxyPluginOptions{
			HTTPUser:     "user",
			HTTPPassword: "password",
		},
	}

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Proxy-Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:password")))
	if !hp.Auth(req) {
		t.Fatal("auth failed")
	}

	req, _ = http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Proxy-Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:password1")))
	if hp.Auth(req) {
		t.Fatal("auth success")
	}
}
