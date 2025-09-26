package orm

import (
	"fmt"
	"testing"
)

func TestRowsToStruct_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{}
	ptrStruct := &struct{}{}
	keyCol := "key"
	valueCol := "value"
	_, err := o.RowsToStruct(ptrStruct, keyCol, valueCol)
	if err != ErrNotImplement {
		t.Errorf("Expected error %v, got %v", ErrNotImplement, err)
	}
}
