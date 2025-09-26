package request

import (
	"fmt"
	"net"
	"testing"
)

func TestResolver_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	resolver := &net.Resolver{}
	r.Resolver(resolver)
	if r.resolver != resolver {
		t.Errorf("resolver not equal")
	}
}
