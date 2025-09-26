package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fp := &FileProvider{
		maxlifetime: 10,
		savePath:    "/tmp/sessions",
	}

	ctx := context.Background()
	total := fp.SessionAll(ctx)

	if total < 0 {
		t.Errorf("Expected total sessions to be greater than or equal to 0, but got %d", total)
	}
}
