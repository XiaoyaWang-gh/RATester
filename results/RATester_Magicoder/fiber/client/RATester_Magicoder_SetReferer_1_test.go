package client

import (
	"fmt"
	"testing"
)

func TestSetReferer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{}
	referer := "https://example.com"
	req.SetReferer(referer)

	if req.referer != referer {
		t.Errorf("Expected referer to be %s, but got %s", referer, req.referer)
	}
}
