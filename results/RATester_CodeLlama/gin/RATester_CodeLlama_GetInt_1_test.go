package gin

import (
	"fmt"
	"testing"
)

func TestGetInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": 1,
		},
	}
	if c.GetInt("key") != 1 {
		t.Errorf("GetInt() = %v, want %v", c.GetInt("key"), 1)
	}
}
