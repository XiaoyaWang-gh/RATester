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
	proxyFunc := func(*http.Request) (*url.URL, error) {
		return &url.URL{
			Scheme: "http",
			Host:   "localhost:8080",
		}, nil
	}
	b.SetProxy(proxyFunc)
	if b.setting.Proxy == nil {
		t.Errorf("Expected proxy function to be set, but it was not")
	}
}
