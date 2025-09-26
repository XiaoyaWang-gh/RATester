package gin

import (
	"fmt"
	"testing"
)

func TestGetStringMapStringSlice_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": map[string][]string{
				"key": []string{"value"},
			},
		},
	}
	smss := c.GetStringMapStringSlice("key")
	if smss == nil {
		t.Errorf("GetStringMapStringSlice() = %v, want %v", smss, map[string][]string{
			"key": []string{"value"},
		})
	}
}
