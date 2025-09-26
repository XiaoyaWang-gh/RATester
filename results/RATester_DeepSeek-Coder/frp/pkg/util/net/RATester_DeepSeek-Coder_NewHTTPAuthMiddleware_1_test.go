package net

import (
	"fmt"
	"testing"
	"time"
)

func TestNewHTTPAuthMiddleware_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	user := "testUser"
	passwd := "testPasswd"
	authMid := NewHTTPAuthMiddleware(user, passwd)

	if authMid.user != user || authMid.passwd != passwd {
		t.Errorf("Expected user and password to be %s and %s, but got %s and %s", user, passwd, authMid.user, authMid.passwd)
	}

	delay := 5 * time.Second
	authMid.SetAuthFailDelay(delay)

	if authMid.authFailDelay != delay {
		t.Errorf("Expected authFailDelay to be %s, but got %s", delay, authMid.authFailDelay)
	}
}
