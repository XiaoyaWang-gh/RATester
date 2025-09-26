package orm

import (
	"fmt"
	"testing"
)

func TestValuesList_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{}
	results := []ParamsList{}
	exprs := []string{"expr1", "expr2"}
	expectedCount := int64(2)
	expectedError := error(nil)

	count, err := o.ValuesList(&results, exprs...)

	if count != expectedCount {
		t.Errorf("Expected count to be %v, but got %v", expectedCount, count)
	}

	if err != expectedError {
		t.Errorf("Expected error to be %v, but got %v", expectedError, err)
	}
}
