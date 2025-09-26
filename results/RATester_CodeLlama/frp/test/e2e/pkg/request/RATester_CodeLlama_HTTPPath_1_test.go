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
	r.HTTPPath("path")
	if r.path != "path" {
		t.Errorf("HTTPPath failed, expect path: %s, but got: %s", "path", r.path)
	}
}
