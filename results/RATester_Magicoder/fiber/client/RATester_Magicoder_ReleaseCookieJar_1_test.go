package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestReleaseCookieJar_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cookieJar := &CookieJar{
		hostCookies: make(map[string][]*fasthttp.Cookie),
	}

	ReleaseCookieJar(cookieJar)

	if len(cookieJar.hostCookies) != 0 {
		t.Errorf("Expected hostCookies to be empty after ReleaseCookieJar, but got %d", len(cookieJar.hostCookies))
	}
}
