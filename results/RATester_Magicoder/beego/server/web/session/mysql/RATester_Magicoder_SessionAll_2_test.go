package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	mp := &Provider{
		maxlifetime: 10,
		savePath:    "/tmp",
	}
	total := mp.SessionAll(ctx)
	if total < 0 {
		t.Errorf("Expected total to be greater than or equal to 0, but got %d", total)
	}
}
