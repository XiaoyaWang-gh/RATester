package fmt

import (
	"fmt"
	"testing"
)

func TestPrintln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	args := []any{"foo", "bar"}
	expected := "foo bar\n"
	actual := ns.Println(args...)
	if actual != expected {
		t.Errorf("expected %q, but actual %q", expected, actual)
	}
}
