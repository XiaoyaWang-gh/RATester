package utils

import (
	"fmt"
	"testing"
)

func TestIfContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	kvs := &SimpleKVs{
		kvs: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	called := false
	kvs.IfContains("key1", func(value interface{}) {
		called = true
		if value != "value1" {
			t.Errorf("Expected value 'value1', but got '%v'", value)
		}
	})

	if !called {
		t.Error("Action function was not called")
	}

	kvs.IfContains("key3", func(value interface{}) {
		t.Error("Action function should not have been called")
	})
}
