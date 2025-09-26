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
	authMid.SetAuthFailDelay(time.Duration(1000))
	if authMid.authFailDelay != time.Duration(1000) {
		t.Errorf("SetAuthFailDelay failed")
	}
}
