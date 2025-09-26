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
	host := "example.com"
	r.HTTPHost(host)
	if r.host != host {
		t.Errorf("Expected host to be %s, but got %s", host, r.host)
	}
}
