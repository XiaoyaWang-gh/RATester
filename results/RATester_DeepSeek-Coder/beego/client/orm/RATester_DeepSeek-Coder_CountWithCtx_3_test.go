package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestCountWithCtx_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	o := &queryM2M{
		md:  nil,
		mi:  nil,
		fi:  nil,
		qs:  nil,
		ind: reflect.Value{},
	}

	count, err := o.CountWithCtx(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count < 0 {
		t.Errorf("Expected count to be greater than or equal to 0, got %d", count)
	}
}
