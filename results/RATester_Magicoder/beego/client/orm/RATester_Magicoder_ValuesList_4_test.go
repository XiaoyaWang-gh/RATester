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

	o := &rawSet{}
	container := []ParamsList{}
	cols := []string{"col1", "col2"}
	expectedCount := int64(1)
	expectedError := error(nil)

	count, err := o.ValuesList(&container, cols...)

	if count != expectedCount {
		t.Errorf("Expected count %d, but got %d", expectedCount, count)
	}

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
}
