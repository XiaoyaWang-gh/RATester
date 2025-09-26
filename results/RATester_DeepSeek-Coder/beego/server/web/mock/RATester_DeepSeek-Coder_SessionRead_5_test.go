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

	t.Parallel()

	ctx := context.Background()
	store := &SessionStore{}
	provider := &SessionProvider{Store: store}

	_, err := provider.SessionRead(ctx, "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
