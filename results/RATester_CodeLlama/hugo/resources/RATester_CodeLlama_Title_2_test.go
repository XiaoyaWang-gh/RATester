package resources

import (
	"fmt"
	"testing"
)

func TestTitle_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &genericResource{
		title: "title",
	}
	if l.Title() != "title" {
		t.Errorf("expected title to be %q but was %q", "title", l.Title())
	}
}
