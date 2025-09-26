package json

import (
	"fmt"
	"testing"
)

func TestInt64_7(t *testing.T) {
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
	v, err := c.Int64("key")
	if err != nil {
		t.Error(err)
	}
	if v != int64(100) {
		t.Errorf("Expected 100, but got %d", v)
	}
}
