package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSetByHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	jar := &CookieJar{}

	host := []byte("example.com")
	cookie1 := &fasthttp.Cookie{}
	cookie1.SetKey("cookie1")
	cookie1.SetValue("value1")
	cookie2 := &fasthttp.Cookie{}
	cookie2.SetKey("cookie2")
	cookie2.SetValue("value2")

	jar.SetByHost(host, cookie1, cookie2)

	hostCookies := jar.getByHostAndPath(host, nil)
	if len(hostCookies) != 2 {
		t.Fatalf("Expected 2 cookies, got %d", len(hostCookies))
	}

	cookie1Found := false
	cookie2Found := false
	for _, cookie := range hostCookies {
		if string(cookie.Key()) == "cookie1" && string(cookie.Value()) == "value1" {
			cookie1Found = true
		}
		if string(cookie.Key()) == "cookie2" && string(cookie.Value()) == "value2" {
			cookie2Found = true
		}
	}

	if !cookie1Found || !cookie2Found {
		t.Fatalf("Expected cookies not found")
	}
}
