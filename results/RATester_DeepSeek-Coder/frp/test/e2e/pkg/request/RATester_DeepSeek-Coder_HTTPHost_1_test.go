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
	host := "test.com"
	r.HTTPHost(host)
	if r.host != host {
		t.Errorf("Expected host to be %s, got %s", host, r.host)
	}
}
