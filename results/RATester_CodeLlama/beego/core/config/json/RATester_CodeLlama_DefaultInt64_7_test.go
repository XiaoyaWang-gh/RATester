package json

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &JSONConfigContainer{
		data: map[string]interface{}{
			"key": int64(100),
		},
	}
	if v := c.DefaultInt64("key", 10); v != 100 {
		t.Fatalf("DefaultInt64 failed")
	}
	if v := c.DefaultInt64("key1", 10); v != 10 {
		t.Fatalf("DefaultInt64 failed")
	}
}
