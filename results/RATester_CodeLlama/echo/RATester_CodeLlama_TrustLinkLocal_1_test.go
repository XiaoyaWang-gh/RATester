package echo

import (
	"fmt"
	"testing"
)

func TestTrustLinkLocal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var v bool
	c := &ipChecker{}
	TrustLinkLocal(v)(c)
	if c.trustLinkLocal != v {
		t.Errorf("TrustLinkLocal(%v) = %v, want %v", v, c.trustLinkLocal, v)
	}
}
