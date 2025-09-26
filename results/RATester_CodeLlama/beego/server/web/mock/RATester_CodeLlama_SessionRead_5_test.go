package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRead_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionProvider{}
	ctx := context.Background()
	sid := "123"
	store, err := s.SessionRead(ctx, sid)
	if err != nil {
		t.Errorf("SessionRead() error = %v", err)
		return
	}
	if store == nil {
		t.Errorf("SessionRead() store = nil")
	}
}
