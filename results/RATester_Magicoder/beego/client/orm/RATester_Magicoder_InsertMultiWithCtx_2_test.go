package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertMultiWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	orm := &DoNothingOrm{}

	bulk := 10
	mds := []interface{}{
		// Add your test data here
	}

	_, err := orm.InsertMultiWithCtx(ctx, bulk, mds)
	if err != nil {
		t.Errorf("InsertMultiWithCtx returned an error: %v", err)
	}
}
