package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestDeleteWithCtx_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	orm := &DoNothingOrm{}

	_, err := orm.DeleteWithCtx(ctx, nil)
	if err != nil {
		t.Errorf("DeleteWithCtx returned an error: %v", err)
	}
}
