package client

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSetKeyValueBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cj := &CookieJar{
		hostCookies: make(map[string][]*fasthttp.Cookie),
	}

	host := "example.com"
	key := []byte("testKey")
	value := []byte("testValue")

	cj.SetKeyValueBytes(host, key, value)

	cookies := cj.getCookiesByHost(host)
	if len(cookies) != 1 {
		t.Fatalf("Expected 1 cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if !bytes.Equal(cookie.Key(), key) {
		t.Errorf("Expected key %s, got %s", key, cookie.Key())
	}

	if !bytes.Equal(cookie.Value(), value) {
		t.Errorf("Expected value %s, got %s", value, cookie.Value())
	}
}
