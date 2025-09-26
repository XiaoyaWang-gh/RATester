package maps

import (
	"fmt"
	"testing"
)

func TestLen_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache[string, string]{
		m: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	}
	if got := c.Len(); got != 3 {
		t.Errorf("Len() = %v, want %v", got, 3)
	}
}
