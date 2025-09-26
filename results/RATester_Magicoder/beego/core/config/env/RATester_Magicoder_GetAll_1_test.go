package env

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGetAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	env := sync.Map{}
	env.Store("key1", "value1")
	env.Store("key2", "value2")

	expected := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	result := GetAll()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
