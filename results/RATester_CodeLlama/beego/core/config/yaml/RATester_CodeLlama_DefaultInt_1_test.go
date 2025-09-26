package yaml

import (
	"fmt"
	"testing"
)

func TestDefaultInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": 1,
		},
	}
	key := "key"
	defaultVal := 2
	v := c.DefaultInt(key, defaultVal)
	if v != 1 {
		t.Errorf("DefaultInt() = %v, want %v", v, 1)
	}
}
