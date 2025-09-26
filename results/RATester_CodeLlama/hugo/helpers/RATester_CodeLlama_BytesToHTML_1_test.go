package helpers

import (
	"fmt"
	"html/template"
	"testing"
)

func TestBytesToHTML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := []byte("hello")
	want := template.HTML("hello")
	got := BytesToHTML(b)
	if got != want {
		t.Errorf("BytesToHTML(%q) = %q; want %q", b, got, want)
	}
}
