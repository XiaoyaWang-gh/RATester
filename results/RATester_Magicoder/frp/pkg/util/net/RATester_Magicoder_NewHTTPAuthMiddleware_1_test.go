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

	if authMid.user != user {
		t.Errorf("Expected user %s, got %s", user, authMid.user)
	}

	if authMid.passwd != passwd {
		t.Errorf("Expected password %s, got %s", passwd, authMid.passwd)
	}

	delay := 5 * time.Second
	authMid.SetAuthFailDelay(delay)

	if authMid.authFailDelay != delay {
		t.Errorf("Expected auth fail delay %s, got %s", delay, authMid.authFailDelay)
	}
}
