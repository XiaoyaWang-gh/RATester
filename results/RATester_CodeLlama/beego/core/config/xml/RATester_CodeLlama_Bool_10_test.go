package xml

import (
	"fmt"
	"testing"
)

func TestBool_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "true",
		},
	}
	b, err := c.Bool("key")
	if err != nil {
		t.Errorf("TestBool error: %v", err)
	}
	if b != true {
		t.Errorf("TestBool error: %v", b)
	}
}
