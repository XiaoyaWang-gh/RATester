package orm

import (
	"fmt"
	"testing"
)

func TestDelete_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	type TestStruct struct {
		ID int64
	}
	testStruct := TestStruct{ID: 1}
	_, err := o.Delete(&testStruct)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
}
