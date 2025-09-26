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

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key": "123",
		},
	}

	val, err := cc.Int("key")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if val != 123 {
		t.Errorf("expected 123, got %d", val)
	}
}
