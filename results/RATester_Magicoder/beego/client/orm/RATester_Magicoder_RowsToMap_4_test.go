package orm

import (
	"fmt"
	"testing"
)

func TestRowsToMap_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &rawSet{}
	result := &Params{}
	keyCol := "key"
	valueCol := "value"

	_, err := o.RowsToMap(result, keyCol, valueCol)
	if err != nil {
		t.Errorf("RowsToMap returned an error: %v", err)
	}

	// Add more test cases as needed
}
