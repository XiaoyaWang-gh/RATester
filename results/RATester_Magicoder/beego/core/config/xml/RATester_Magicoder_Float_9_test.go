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

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key": "123.456",
		},
	}

	floatVal, err := cc.Float("key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if floatVal != 123.456 {
		t.Errorf("Expected 123.456, but got %v", floatVal)
	}
}
