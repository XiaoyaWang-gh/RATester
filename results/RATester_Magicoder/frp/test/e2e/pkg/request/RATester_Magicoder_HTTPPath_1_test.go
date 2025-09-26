package request

import (
	"fmt"
	"testing"
)

func TestHTTPPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	path := "/test"
	r.HTTPPath(path)
	if r.path != path {
		t.Errorf("Expected path to be %s, but got %s", path, r.path)
	}
}
