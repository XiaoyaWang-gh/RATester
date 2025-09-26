package config

import (
	"fmt"
	"testing"
)

func TestOnChange_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testKey := "testKey"
	testValue := "testValue"

	testFn := func(value string) {
		if value != testValue {
			t.Errorf("Expected value %s, got %s", testValue, value)
		}
	}

	OnChange(testKey, testFn)

	globalInstance.Set(testKey, testValue)
}
