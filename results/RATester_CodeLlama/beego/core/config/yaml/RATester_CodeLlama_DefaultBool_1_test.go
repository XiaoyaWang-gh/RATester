package yaml

import (
	"fmt"
	"testing"
)

func TestDefaultBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "true",
		},
	}
	if v := c.DefaultBool("key", false); v != true {
		t.Errorf("DefaultBool() = %v, want %v", v, true)
	}
}
