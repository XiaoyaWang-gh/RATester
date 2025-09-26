package client

import (
	"fmt"
	"testing"
)

func TestHeader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	r.RawResponse.Header.Set("key", "value")
	if r.Header("key") != "value" {
		t.Errorf("Header() = %v, want %v", r.Header("key"), "value")
	}
}
