package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestUpdate_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var values Params
	var result int64
	var err error

	// TEST CASE 1
	result, err = o.Update(values)
	if err != nil {
		t.Errorf("TestUpdate failed, err: %v", err)
	}
	if result != 0 {
		t.Errorf("TestUpdate failed, result: %v", result)
	}

	// TEST CASE 2
	result, err = o.UpdateWithCtx(context.Background(), values)
	if err != nil {
		t.Errorf("TestUpdate failed, err: %v", err)
	}
	if result != 0 {
		t.Errorf("TestUpdate failed, result: %v", result)
	}
}
