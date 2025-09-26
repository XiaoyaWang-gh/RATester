package maps

import (
	"fmt"
	"testing"
)

func TestGetOrCreate_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache[string, int]{m: map[string]int{}}
	v, err := c.GetOrCreate("foo", func() (int, error) {
		return 1, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if v != 1 {
		t.Errorf("expected 1, got %d", v)
	}
	if c.Len() != 1 {
		t.Errorf("expected 1, got %d", c.Len())
	}
	v, err = c.GetOrCreate("foo", func() (int, error) {
		return 2, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if v != 1 {
		t.Errorf("expected 1, got %d", v)
	}
	if c.Len() != 1 {
		t.Errorf("expected 1, got %d", c.Len())
	}
}
