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

	pder := &CookieProvider{}
	ctx := context.Background()
	sid := "123"
	err := pder.SessionUpdate(ctx, sid)
	if err != nil {
		t.Errorf("SessionUpdate() error = %v, want nil", err)
		return
	}
}
