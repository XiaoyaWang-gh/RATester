package yaml

import (
	"fmt"
	"testing"
)

func TestOnChange_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: make(map[string]interface{}),
	}

	cc.OnChange("test", func(value string) {
		t.Log("OnChange called with value:", value)
	})

	// Add assertions here to verify that OnChange was called as expected
}
