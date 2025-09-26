package gin

import (
	"fmt"
	"testing"
)

func TestGetStringSlice_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": []string{"value1", "value2"},
		},
	}
	ss := c.GetStringSlice("key")
	if len(ss) != 2 {
		t.Errorf("GetStringSlice() = %v, want 2", len(ss))
	}
	if ss[0] != "value1" {
		t.Errorf("GetStringSlice() = %v, want value1", ss[0])
	}
	if ss[1] != "value2" {
		t.Errorf("GetStringSlice() = %v, want value2", ss[1])
	}
}
