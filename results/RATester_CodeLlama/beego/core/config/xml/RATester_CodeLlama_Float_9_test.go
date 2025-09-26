package xml

import (
	"fmt"
	"testing"
)

func TestFloat_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "123.456",
		},
	}
	v, err := c.Float("key")
	if err != nil {
		t.Error(err)
	}
	if v != 123.456 {
		t.Errorf("expect 123.456, but got %v", v)
	}
}
