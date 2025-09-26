package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestDeleteWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}
	_, err := d.DeleteWithCtx(ctx)
	if err != nil {
		t.Errorf("DeleteWithCtx() error = %v, wantErr %v", err, nil)
		return
	}
}
