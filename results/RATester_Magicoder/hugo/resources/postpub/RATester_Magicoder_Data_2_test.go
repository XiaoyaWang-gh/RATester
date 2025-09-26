package postpub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestData_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{
		prefix: "test",
	}

	expected := map[string]any{
		"Integrity": "",
	}
	insertFieldPlaceholders("Data", expected, r.field)

	result := r.Data()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
