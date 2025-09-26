package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestExistWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQueryM2Mer{}
	ctx := context.Background()
	i := 1
	if d.ExistWithCtx(ctx, i) != true {
		t.Errorf("ExistWithCtx() = %v, want %v", d.ExistWithCtx(ctx, i), true)
	}
}
