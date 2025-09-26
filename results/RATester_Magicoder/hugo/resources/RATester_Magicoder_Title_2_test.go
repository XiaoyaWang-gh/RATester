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
		title: "Test Title",
	}

	expected := "Test Title"
	actual := l.Title()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
