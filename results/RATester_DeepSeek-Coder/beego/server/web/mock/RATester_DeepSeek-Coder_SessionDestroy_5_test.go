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
	s := &SessionProvider{
		Store: &SessionStore{
			sid:    "test",
			values: make(map[interface{}]interface{}),
		},
	}

	s.Store.values["key"] = "value"

	err := s.SessionDestroy(ctx, "test")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if s.Store != nil {
		t.Errorf("Expected Store to be nil, got %v", s.Store)
	}
}
