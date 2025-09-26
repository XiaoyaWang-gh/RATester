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

	o := &rawSet{
		query: "SELECT * FROM test",
		args:  []interface{}{},
		orm:   &ormBase{},
	}

	result := &Params{}
	keyCol := "key"
	valueCol := "value"

	count, err := o.RowsToMap(result, keyCol, valueCol)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count == 0 {
		t.Errorf("Expected count to be greater than 0, got %v", count)
	}

	// Add more specific tests here
}
