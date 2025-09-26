package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	sp := &SessionProvider{
		Store: &SessionStore{
			sid:    "test",
			values: make(map[interface{}]interface{}),
		},
	}
	result := sp.SessionAll(ctx)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
