package session

import (
	"fmt"
	"testing"
)

func TestLen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &data{
		Data: map[string]any{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}
	if d.Len() != 3 {
		t.Errorf("expected 3, got %d", d.Len())
	}
}
