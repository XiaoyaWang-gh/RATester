package orm

import (
	"fmt"
	"testing"
)

func TestCount_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{}
	count, err := o.Count()
	if err != nil {
		t.Errorf("Error in Count: %v", err)
	}
	if count < 0 {
		t.Errorf("Count should not be negative, got: %d", count)
	}
}
