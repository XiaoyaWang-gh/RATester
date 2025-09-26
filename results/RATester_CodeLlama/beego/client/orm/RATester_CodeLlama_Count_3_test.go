package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestCount_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var err error
	var count int64

	// test case 1:
	count, err = o.Count()
	if err != nil {
		t.Errorf("TestCount error, err=%v", err)
		return
	}
	if count != 0 {
		t.Errorf("TestCount error, count=%v", count)
		return
	}

	// test case 2:
	count, err = o.CountWithCtx(context.Background())
	if err != nil {
		t.Errorf("TestCount error, err=%v", err)
		return
	}
	if count != 0 {
		t.Errorf("TestCount error, count=%v", count)
		return
	}
}
