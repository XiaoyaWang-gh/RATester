package gin

import (
	"fmt"
	"testing"
)

func TestGetFloat64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": 1.23,
		},
	}
	f64 := c.GetFloat64("key")
	if f64 != 1.23 {
		t.Errorf("GetFloat64() = %v, want %v", f64, 1.23)
	}
}
