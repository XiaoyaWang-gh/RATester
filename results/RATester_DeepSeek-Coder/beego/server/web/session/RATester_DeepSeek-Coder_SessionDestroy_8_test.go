package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &CookieProvider{}

	err := pder.SessionDestroy(ctx, "test_sid")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
