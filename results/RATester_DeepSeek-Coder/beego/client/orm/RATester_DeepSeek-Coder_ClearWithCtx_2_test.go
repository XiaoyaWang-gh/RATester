package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestClearWithCtx_2(t *testing.T) {
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

	_, err := o.ClearWithCtx(ctx)
	if err != nil {
		t.Errorf("ClearWithCtx() error = %v, wantErr %v", err, false)
		return
	}
}
