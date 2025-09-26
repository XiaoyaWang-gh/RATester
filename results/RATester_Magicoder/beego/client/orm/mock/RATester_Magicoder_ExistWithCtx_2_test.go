package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestExistWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}

	result := d.ExistWithCtx(ctx)

	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}
