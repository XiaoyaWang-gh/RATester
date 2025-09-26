package fmt

import (
	"fmt"
	"testing"
)

func TestPrint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	args := []any{"foo", "bar"}
	got := ns.Print(args...)
	want := "foo bar"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
