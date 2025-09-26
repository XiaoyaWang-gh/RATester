package request

import (
	"fmt"
	"testing"
)

func TestHTTPHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.HTTPHost("www.example.com")
	if r.host != "www.example.com" {
		t.Errorf("HTTPHost failed")
	}
}
