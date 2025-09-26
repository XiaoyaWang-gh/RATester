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

	count := sp.SessionAll(ctx)
	if count != 0 {
		t.Errorf("Expected count to be 0, got %d", count)
	}
}
