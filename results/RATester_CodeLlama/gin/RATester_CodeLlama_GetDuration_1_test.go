package gin

import (
	"fmt"
	"testing"
)

func TestGetDuration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": 100,
		},
	}
	d := c.GetDuration("key")
	if d != 100 {
		t.Errorf("GetDuration() = %v, want %v", d, 100)
	}
}
