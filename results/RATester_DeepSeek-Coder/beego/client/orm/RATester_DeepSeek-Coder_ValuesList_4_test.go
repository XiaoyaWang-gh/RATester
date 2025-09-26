package orm

import (
	"fmt"
	"testing"
)

func TestValuesList_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &rawSet{
		query: "SELECT * FROM table",
		args:  []interface{}{},
	}

	container := &[]ParamsList{}
	expectedCount := 10
	expectedError := error(nil)

	count, err := o.ValuesList(container, "col1", "col2")

	if count != int64(expectedCount) {
		t.Errorf("Expected count to be %v, got %v", expectedCount, count)
	}

	if err != expectedError {
		t.Errorf("Expected error to be %v, got %v", expectedError, err)
	}
}
