package httplib

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestSetProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}
	proxy := func(req *http.Request) (*url.URL, error) {
		return &url.URL{
			Scheme: "http",
			Host:   "proxy.example.com",
		}, nil
	}
	b.SetProxy(proxy)
	if b.setting.Proxy == nil {
		t.Error("Proxy not set")
	}
}
