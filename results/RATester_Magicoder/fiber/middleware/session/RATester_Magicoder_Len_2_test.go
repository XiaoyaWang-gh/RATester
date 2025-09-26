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
		},
	}

	expected := 2
	actual := d.Len()

	if actual != expected {
		t.Errorf("Expected Len() to return %d, but got %d", expected, actual)
	}
}
