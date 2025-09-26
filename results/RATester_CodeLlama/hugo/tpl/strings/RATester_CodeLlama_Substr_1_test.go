package strings

import (
	"fmt"
	"testing"
)

func TestSubstr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	s, err := ns.Substr("hello", 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if s != "el" {
		t.Errorf("got %q", s)
	}
}
