package json

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": "value1;value2;value3",
		},
	}

	result, err := container.Strings("key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []string{"value1", "value2", "value3"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
