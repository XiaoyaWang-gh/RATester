package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	sp := &SessionProvider{}
	err := sp.SessionInit(ctx, 10, "config")
	if err != nil {
		t.Errorf("SessionInit failed: %v", err)
	}
}
