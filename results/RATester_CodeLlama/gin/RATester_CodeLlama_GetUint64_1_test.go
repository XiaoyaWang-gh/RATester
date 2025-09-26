package gin

import (
	"fmt"
	"testing"
)

func TestGetUint64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": uint64(1),
		},
	}
	if ui64 := c.GetUint64("key"); ui64 != 1 {
		t.Errorf("GetUint64() = %v, want %v", ui64, 1)
	}
}
