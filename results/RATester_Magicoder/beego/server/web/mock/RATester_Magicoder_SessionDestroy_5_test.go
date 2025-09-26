package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	sp := &SessionProvider{
		Store: &SessionStore{
			sid:    "test_sid",
			values: make(map[interface{}]interface{}),
		},
	}

	err := sp.SessionDestroy(ctx, "test_sid")
	if err != nil {
		t.Errorf("SessionDestroy returned an error: %v", err)
	}

	if sp.Store != nil {
		t.Errorf("SessionDestroy did not destroy the session store")
	}
}
