package hugofs

import (
	"fmt"
	"testing"
)

func TestClean_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rm := &RootMapping{
		From: "  /foo/bar/",
		To:   "/foo/bar",
	}
	rm.clean()
	if rm.From != "/foo/bar" {
		t.Errorf("From should be cleaned: %q", rm.From)
	}
	if rm.To != "/foo/bar" {
		t.Errorf("To should be cleaned: %q", rm.To)
	}
}
