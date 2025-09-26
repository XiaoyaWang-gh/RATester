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

	key := "testKey"
	expectedValue := "testValue"

	OnChange(key, func(value string) {
		if value != expectedValue {
			t.Errorf("Expected value %s, but got %s", expectedValue, value)
		}
	})

	globalInstance.Set(key, expectedValue)
}
