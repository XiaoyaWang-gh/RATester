package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pder := &CookieProvider{}
	ctx := context.Background()
	if pder.SessionAll(ctx) != 0 {
		t.Error("SessionAll() error")
	}
}
