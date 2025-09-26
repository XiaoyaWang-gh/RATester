package maps

import (
	"fmt"
	"testing"
)

func TestDelete_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Scratch{
		values: map[string]any{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
	}
	key := "key2"
	want := "value2"
	got := c.Delete(key)
	if got != want {
		t.Errorf("Delete() = %v, want %v", got, want)
	}
	if len(c.values) != 2 {
		t.Errorf("Delete() = %v, want %v", len(c.values), 2)
	}
}
