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
	sp := &SessionProvider{}

	exist, err := sp.SessionExist(ctx, "test_sid")
	if err != nil {
		t.Errorf("SessionExist() error = %v, wantErr %v", err, false)
		return
	}
	if !exist {
		t.Errorf("SessionExist() = %v, want %v", exist, true)
	}
}
