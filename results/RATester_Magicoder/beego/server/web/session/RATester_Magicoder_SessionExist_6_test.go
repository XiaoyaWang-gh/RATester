package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionExist_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &CookieProvider{}

	sid := "test_session_id"
	exist, err := pder.SessionExist(ctx, sid)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !exist {
		t.Errorf("Expected session to exist, but it does not")
	}
}
