package client

import (
	"fmt"
	"testing"
)

func TestSetCookieJar_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	cookieJar := &CookieJar{}
	client.SetCookieJar(cookieJar)

	if client.cookieJar != cookieJar {
		t.Errorf("Expected cookieJar to be %v, but got %v", cookieJar, client.cookieJar)
	}
}
