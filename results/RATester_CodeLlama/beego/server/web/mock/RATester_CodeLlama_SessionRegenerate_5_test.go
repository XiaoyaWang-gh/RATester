package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionProvider{}
	oldsid := "oldsid"
	sid := "sid"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "oldsid", oldsid)
	ctx = context.WithValue(ctx, "sid", sid)
	store, err := s.SessionRegenerate(ctx, oldsid, sid)
	if err != nil {
		t.Errorf("SessionRegenerate error %v", err)
	}
	if store == nil {
		t.Errorf("SessionRegenerate store is nil")
	}
}
