package xml

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": 1.2,
		},
	}
	if v := c.DefaultFloat("key", 1.3); v != 1.2 {
		t.Errorf("DefaultFloat() = %v, want %v", v, 1.2)
	}
	if v := c.DefaultFloat("key1", 1.3); v != 1.3 {
		t.Errorf("DefaultFloat() = %v, want %v", v, 1.3)
	}
}
