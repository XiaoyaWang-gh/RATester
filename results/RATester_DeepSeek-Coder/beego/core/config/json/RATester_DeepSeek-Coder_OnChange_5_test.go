package json

import (
	"fmt"
	"testing"
)

func TestOnChange_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	jcc := &JSONConfigContainer{
		data: make(map[string]interface{}),
	}

	key := "testKey"
	expectedValue := "testValue"

	jcc.data[key] = expectedValue

	called := false
	fn := func(value string) {
		called = true
		if value != expectedValue {
			t.Errorf("Expected value to be %s, got %s", expectedValue, value)
		}
	}

	jcc.OnChange(key, fn)

	if !called {
		t.Errorf("Expected function to be called, but it wasn't")
	}
}
