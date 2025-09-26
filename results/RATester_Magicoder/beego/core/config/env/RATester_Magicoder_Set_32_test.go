package env

import (
	"fmt"
	"testing"
)

func TestSet_32(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "testKey"
	value := "testValue"
	Set(key, value)

	storedValue, ok := env.Load(key)
	if !ok {
		t.Errorf("Expected key %s to be in the map", key)
	}

	if storedValue != value {
		t.Errorf("Expected value for key %s to be %s, but got %s", key, value, storedValue)
	}
}
