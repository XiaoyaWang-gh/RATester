package fmt

import (
	"fmt"
	"testing"
)

func TestPrintf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	format := "Hello, %s!"
	args := []any{"world"}
	got := ns.Printf(format, args...)
	want := "Hello, world!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
