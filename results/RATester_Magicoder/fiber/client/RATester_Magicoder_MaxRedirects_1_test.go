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

	req := &Request{
		maxRedirects: 5,
	}

	if got := req.MaxRedirects(); got != 5 {
		t.Errorf("MaxRedirects() = %v, want %v", got, 5)
	}
}
