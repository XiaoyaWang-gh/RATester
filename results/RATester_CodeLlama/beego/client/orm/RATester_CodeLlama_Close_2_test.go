package orm

import (
	"fmt"
	"testing"
)

func TestClose_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &rawPrepare{}
	if err := o.Close(); err != nil {
		t.Errorf("Close() error = %v", err)
		return
	}
	if o.closed != true {
		t.Errorf("Close() o.closed = %v, want %v", o.closed, true)
	}
}
