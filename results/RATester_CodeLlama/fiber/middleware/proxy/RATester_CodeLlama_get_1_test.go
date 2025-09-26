package proxy

import (
	"fmt"
	"testing"
)

func TestGet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &roundrobin{
		pool: []string{"a", "b", "c"},
	}
	for i := 0; i < 100; i++ {
		if r.get() != "a" {
			t.Errorf("get() = %v, want %v", r.get(), "a")
		}
	}
}
