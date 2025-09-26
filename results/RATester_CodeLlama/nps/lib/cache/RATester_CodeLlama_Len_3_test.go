package cache

import (
	"fmt"
	"testing"
)

func TestLen_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := New(10)
	for i := 0; i < 10; i++ {
		c.Add(i, i)
	}
	if c.Len() != 10 {
		t.Fatalf("expect 10, but %d", c.Len())
	}
}
