package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestNewSessionProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	sp := newSessionProvider()

	if sp == nil {
		t.Errorf("newSessionProvider() = %v, want not nil", sp)
	}

	if sp.Store == nil {
		t.Errorf("newSessionProvider().Store = %v, want not nil", sp.Store)
	}

	_, err := sp.SessionRead(ctx, "test")
	if err != nil {
		t.Errorf("newSessionProvider().SessionRead() error = %v, want nil", err)
	}
}
