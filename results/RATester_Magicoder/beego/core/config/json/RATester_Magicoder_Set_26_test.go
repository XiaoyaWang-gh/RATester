package json

import (
	"fmt"
	"testing"
)

func TestSet_26(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: make(map[string]interface{}),
	}

	key := "testKey"
	val := "testValue"

	err := container.Set(key, val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if container.data[key] != val {
		t.Errorf("Expected value %s for key %s, but got %v", val, key, container.data[key])
	}
}
