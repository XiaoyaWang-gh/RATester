package echo

import (
	"fmt"
	"testing"
)

func TestURI_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	handler := func(c Context) error {
		return nil
	}
	expected := "/test"
	e.Router().Add("GET", "/test", handler)
	actual := e.URI(handler)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
