package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := make(unsafeHeader)
	h.Set("key1", "value1")
	h.Set("key1", "value2")
	h.Set("key2", "value3")

	values := h.Values("key1")
	if len(values) != 2 || values[0] != "value1" || values[1] != "value2" {
		t.Errorf("Expected values for key1 to be [value1, value2], but got %v", values)
	}

	values = h.Values("key2")
	if len(values) != 1 || values[0] != "value3" {
		t.Errorf("Expected values for key2 to be [value3], but got %v", values)
	}

	values = h.Values("key3")
	if len(values) != 0 {
		t.Errorf("Expected values for non-existent key to be [], but got %v", values)
	}
}
