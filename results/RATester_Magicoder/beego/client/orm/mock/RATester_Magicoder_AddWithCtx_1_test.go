package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestAddWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQueryM2Mer{}
	_, err := d.AddWithCtx(ctx, "test")
	if err != nil {
		t.Errorf("AddWithCtx() error = %v, wantErr %v", err, false)
		return
	}
}
