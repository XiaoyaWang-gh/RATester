package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSet_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	jar := &CookieJar{
		hostCookies: make(map[string][]*fasthttp.Cookie),
	}

	uri := &fasthttp.URI{}
	uri.Parse(nil, []byte("http://example.com"))

	cookie := &fasthttp.Cookie{}
	cookie.SetKey("test")
	cookie.SetValue("value")

	jar.Set(uri, cookie)

	cookies := jar.getCookiesByHost("example.com")
	if len(cookies) != 1 {
		t.Fatalf("Expected 1 cookie, got %d", len(cookies))
	}

	if string(cookies[0].Key()) != "test" {
		t.Fatalf("Expected cookie key to be 'test', got %s", string(cookies[0].Key()))
	}

	if string(cookies[0].Value()) != "value" {
		t.Fatalf("Expected cookie value to be 'value', got %s", string(cookies[0].Value()))
	}
}
