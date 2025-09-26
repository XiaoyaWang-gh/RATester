package yaml

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": float64(1.1),
		},
	}
	v := c.DefaultFloat("key", 2.2)
	if v != 1.1 {
		t.Errorf("DefaultFloat() = %v, want %v", v, 1.1)
	}

	v = c.DefaultFloat("key1", 2.2)
	if v != 2.2 {
		t.Errorf("DefaultFloat() = %v, want %v", v, 2.2)
	}
}
