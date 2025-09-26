package xml

import (
	"fmt"
	"testing"
)

func TestOnChange_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: make(map[string]interface{}),
	}

	key := "testKey"
	expectedValue := "testValue"

	cc.OnChange(key, func(value string) {
		if value != expectedValue {
			t.Errorf("Expected value %s, but got %s", expectedValue, value)
		}
	})

	cc.Set(key, expectedValue)
}
