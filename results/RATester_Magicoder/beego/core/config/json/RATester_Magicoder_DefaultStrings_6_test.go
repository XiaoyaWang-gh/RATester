package json

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": []string{"value1", "value2"},
		},
	}

	defaultVal := []string{"default1", "default2"}
	result := container.DefaultStrings("key", defaultVal)

	if !reflect.DeepEqual(result, []string{"value1", "value2"}) {
		t.Errorf("Expected %v, got %v", []string{"value1", "value2"}, result)
	}

	result = container.DefaultStrings("nonExistentKey", defaultVal)
	if !reflect.DeepEqual(result, defaultVal) {
		t.Errorf("Expected %v, got %v", defaultVal, result)
	}
}
