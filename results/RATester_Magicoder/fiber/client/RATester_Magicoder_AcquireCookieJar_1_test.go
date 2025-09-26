package client

import (
	"fmt"
	"testing"
)

func TestAcquireCookieJar_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	jar := AcquireCookieJar()

	if jar == nil {
		t.Error("Expected a CookieJar, but got nil")
	}

	jar.Release()
}
