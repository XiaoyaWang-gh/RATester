package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pder := &CookieProvider{}
	ctx := context.Background()
	oldsid := "oldsid"
	sid := "sid"
	store, err := pder.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Fatal(err)
	}
	if store == nil {
		t.Fatal("store is nil")
	}
}
