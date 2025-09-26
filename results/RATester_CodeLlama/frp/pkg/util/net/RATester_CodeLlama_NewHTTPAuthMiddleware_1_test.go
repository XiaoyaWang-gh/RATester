package net

import (
	"fmt"
	"testing"
)

func TestNewHTTPAuthMiddleware_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	authMid := NewHTTPAuthMiddleware("user", "passwd")
	if authMid.user != "user" {
		t.Error("NewHTTPAuthMiddleware failed")
	}
}
