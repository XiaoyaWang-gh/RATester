package xml

import (
	"fmt"
	"testing"
	"time"
)

func TestOnChange_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	called := false
	cc.OnChange("key1", func(value string) {
		called = true
	})

	if called {
		t.Error("OnChange should not be called for non-existing key")
	}

	cc.data["key1"] = "newValue"

	time.Sleep(1 * time.Second)

	if !called {
		t.Error("OnChange should be called for existing key")
	}
}
