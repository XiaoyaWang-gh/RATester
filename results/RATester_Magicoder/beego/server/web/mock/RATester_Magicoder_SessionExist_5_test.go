package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionExist_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	mockProvider := &SessionProvider{}

	exist, err := mockProvider.SessionExist(ctx, "test_sid")
	if err != nil {
		t.Errorf("SessionExist returned an error: %v", err)
	}
	if !exist {
		t.Error("SessionExist returned false, expected true")
	}
}
