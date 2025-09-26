package yaml

import (
	"fmt"
	"testing"
)

func TestInt64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": int64(100),
		},
	}
	v, err := c.Int64("key")
	if err != nil {
		t.Errorf("Int64() error = %v", err)
		return
	}
	if v != int64(100) {
		t.Errorf("Int64() = %v, want %v", v, int64(100))
	}
}
