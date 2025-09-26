package request

import (
	"fmt"
	"testing"
)

func TestHTTP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.HTTP()
	if r.protocol != "http" {
		t.Errorf("Expected protocol to be 'http', but got '%s'", r.protocol)
	}
}
