package gin

import (
	"fmt"
	"testing"
)

func TestGetInt64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": int64(1),
		},
	}
	if c.GetInt64("key") != int64(1) {
		t.Errorf("GetInt64() = %v, want %v", c.GetInt64("key"), int64(1))
	}
}
