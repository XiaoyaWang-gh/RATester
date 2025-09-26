package hashing

import (
	"fmt"
	"testing"
)

func TestHashString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vs := []any{1, "2", 3.0}
	want := "123"
	got := HashString(vs...)
	if got != want {
		t.Errorf("HashString(%v) = %v; want %v", vs, got, want)
	}
}
