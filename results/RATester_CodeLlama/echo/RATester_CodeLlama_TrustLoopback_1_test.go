package echo

import (
	"fmt"
	"testing"
)

func TestTrustLoopback_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var v bool
	c := &ipChecker{}
	TrustLoopback(v)(c)
	if c.trustLoopback != v {
		t.Errorf("TrustLoopback(%v) = %v, want %v", v, c.trustLoopback, v)
	}
}
