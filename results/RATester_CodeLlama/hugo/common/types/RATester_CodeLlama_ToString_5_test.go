package types

import (
	"fmt"
	"testing"
)

func TestToString_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := 1
	want := "1"
	got := ToString(v)
	if got != want {
		t.Errorf("ToString(%v) = %v; want %v", v, got, want)
	}
}
