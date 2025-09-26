package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key": "value1;value2;value3",
		},
	}

	result, err := cc.Strings("key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []string{"value1", "value2", "value3"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
