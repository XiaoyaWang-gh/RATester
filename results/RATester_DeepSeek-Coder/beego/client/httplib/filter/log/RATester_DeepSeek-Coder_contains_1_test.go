package log

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice := []string{"apple", "banana", "cherry"}
	want := true
	got := contains(slice, "banana")
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
