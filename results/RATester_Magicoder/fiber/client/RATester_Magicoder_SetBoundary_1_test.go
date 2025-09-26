package client

import (
	"fmt"
	"testing"
)

func TestSetBoundary_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	boundary := "testBoundary"
	r.SetBoundary(boundary)

	if r.boundary != boundary {
		t.Errorf("Expected boundary to be %s, but got %s", boundary, r.boundary)
	}
}
