package orm

import (
	"fmt"
	"testing"
)

func TestRowsToStruct_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Key   string
		Value string
	}

	testStruct := TestStruct{}
	rawSet := rawSet{
		query: "SELECT key, value FROM table",
		args:  []interface{}{},
	}

	_, err := rawSet.RowsToStruct(&testStruct, "key", "value")
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	// Add assertions here to check if the testStruct is correctly populated
}
