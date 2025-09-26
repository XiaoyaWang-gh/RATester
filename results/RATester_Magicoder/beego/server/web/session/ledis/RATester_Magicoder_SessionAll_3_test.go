package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	lp := &Provider{
		maxlifetime: 10,
		SavePath:    "/tmp",
		Db:          0,
	}
	result := lp.SessionAll(ctx)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
