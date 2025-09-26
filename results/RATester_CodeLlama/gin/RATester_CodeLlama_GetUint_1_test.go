package gin

import (
	"fmt"
	"testing"
)

func TestGetUint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": uint(1),
		},
	}
	if c.GetUint("key") != 1 {
		t.Errorf("GetUint() = %v, want %v", c.GetUint("key"), 1)
	}
}
