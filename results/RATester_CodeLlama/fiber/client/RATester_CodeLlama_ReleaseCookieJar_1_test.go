package client

import (
	"fmt"
	"testing"
)

func TestReleaseCookieJar_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &CookieJar{}
	ReleaseCookieJar(c)
	if c.hostCookies != nil {
		t.Errorf("ReleaseCookieJar() failed")
	}
}
