package strings

import (
	"fmt"
	"testing"
)

func TestContainsAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	ns := &Namespace{}
	var s any
	var chars any
	var want bool
	var got bool
	var err error
	{
		s = "hello"
		chars = "aeiou"
		want = true
		got, err = ns.ContainsAny(s, chars)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	{
		s = "hello"
		chars = "xyz"
		want = false
		got, err = ns.ContainsAny(s, chars)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
