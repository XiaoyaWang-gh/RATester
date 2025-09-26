package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRead_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pder := &CookieProvider{}
	ctx := context.Background()
	sid := "123456"
	rs, err := pder.SessionRead(ctx, sid)
	if err != nil {
		t.Error(err)
	}
	if rs == nil {
		t.Error("rs is nil")
	}
}
