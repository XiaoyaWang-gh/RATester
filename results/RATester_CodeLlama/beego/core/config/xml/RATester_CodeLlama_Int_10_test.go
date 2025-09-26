package xml

import (
	"fmt"
	"testing"
)

func TestInt_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "123",
		},
	}
	v, err := c.Int("key")
	if err != nil {
		t.Error(err)
	}
	if v != 123 {
		t.Errorf("Expected 123, but got %d", v)
	}
}
