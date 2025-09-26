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

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "value",
		},
	}
	c.OnChange("key", func(value string) {
		t.Log("OnChange")
	})
	c.Set("key", "new value")
}
