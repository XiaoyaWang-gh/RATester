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

	pder := &CookieProvider{}
	ctx := context.Background()
	sid := "123"
	exist, err := pder.SessionExist(ctx, sid)
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Error("SessionExist failed")
	}
}
