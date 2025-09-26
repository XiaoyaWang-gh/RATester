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

	cc.OnChange("key", func(value string) {
		t.Logf("Key changed: %s, new value: %s", "key", value)
	})

	cc.Set("key", "newValue")
}
