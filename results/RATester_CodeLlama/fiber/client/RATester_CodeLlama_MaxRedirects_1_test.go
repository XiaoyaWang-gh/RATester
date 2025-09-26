package client

import (
	"fmt"
	"testing"
)

func TestMaxRedirects_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	if r.MaxRedirects() != 0 {
		t.Errorf("MaxRedirects() = %v, want %v", r.MaxRedirects(), 0)
	}
}
