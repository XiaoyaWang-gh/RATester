package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSetKeyValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cj := &CookieJar{
		hostCookies: make(map[string][]*fasthttp.Cookie),
	}

	host := "example.com"
	key := "testKey"
	value := "testValue"

	cj.SetKeyValue(host, key, value)

	cookies := cj.getCookiesByHost(host)
	if len(cookies) != 1 {
		t.Fatalf("Expected 1 cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if string(cookie.Key()) != key {
		t.Errorf("Expected key %s, got %s", key, string(cookie.Key()))
	}

	if string(cookie.Value()) != value {
		t.Errorf("Expected value %s, got %s", value, string(cookie.Value()))
	}
}
