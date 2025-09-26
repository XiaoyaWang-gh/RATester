package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertWithCtx_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	orm := &DoNothingOrm{}

	_, err := orm.InsertWithCtx(ctx, nil)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
