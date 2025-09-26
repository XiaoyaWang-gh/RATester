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

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": "value",
		},
	}

	called := false
	container.OnChange("key", func(value string) {
		called = true
	})

	if !called {
		t.Error("OnChange function was not called")
	}
}
