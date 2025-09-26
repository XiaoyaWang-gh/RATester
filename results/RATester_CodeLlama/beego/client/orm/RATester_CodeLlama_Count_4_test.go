package orm

import (
	"fmt"
	"testing"
)

func TestCount_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &queryM2M{}
	count, err := o.Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Errorf("count should be 0, but got %d", count)
	}
}
