package xml

import (
	"fmt"
	"testing"
)

func TestInt64_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "1234567890",
		},
	}
	v, err := c.Int64("key")
	if err != nil {
		t.Errorf("Int64() error = %v", err)
		return
	}
	if v != 1234567890 {
		t.Errorf("Int64() = %v, want %v", v, 1234567890)
	}
}
