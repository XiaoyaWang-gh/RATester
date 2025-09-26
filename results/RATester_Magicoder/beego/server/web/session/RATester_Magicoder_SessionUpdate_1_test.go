package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionUpdate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &CookieProvider{}

	err := pder.SessionUpdate(ctx, "test_sid")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
