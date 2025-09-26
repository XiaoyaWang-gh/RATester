package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestAllWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}
	container := make([]interface{}, 0)
	cols := []string{"col1", "col2"}
	count, err := d.AllWithCtx(ctx, &container, cols...)
	if err != nil {
		t.Errorf("AllWithCtx returned an error: %v", err)
	}
	if count != 0 {
		t.Errorf("AllWithCtx returned an incorrect count: %d", count)
	}
}
