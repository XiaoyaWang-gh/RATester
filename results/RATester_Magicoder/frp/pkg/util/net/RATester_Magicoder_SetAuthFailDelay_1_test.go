package net

import (
	"fmt"
	"testing"
	"time"
)

func TestSetAuthFailDelay_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	authMid := &HTTPAuthMiddleware{}
	delay := time.Duration(10) * time.Second
	authMid.SetAuthFailDelay(delay)
	if authMid.authFailDelay != delay {
		t.Errorf("Expected authFailDelay to be %v, but got %v", delay, authMid.authFailDelay)
	}
}
