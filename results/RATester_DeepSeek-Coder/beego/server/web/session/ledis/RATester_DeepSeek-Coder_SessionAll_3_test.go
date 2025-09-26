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
		maxlifetime: 100,
		SavePath:    "/tmp",
		Db:          0,
	}

	result := lp.SessionAll(ctx)
	expected := 0

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
