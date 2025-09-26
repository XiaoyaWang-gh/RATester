package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestPrepareInsertWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}
	_, err := d.PrepareInsertWithCtx(ctx)
	if err != nil {
		t.Errorf("PrepareInsertWithCtx() error = %v, wantErr %v", err, false)
		return
	}
}
