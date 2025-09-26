package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pder := &CookieProvider{}

	// Test if the function panics
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	pder.SessionGC(ctx)
}
