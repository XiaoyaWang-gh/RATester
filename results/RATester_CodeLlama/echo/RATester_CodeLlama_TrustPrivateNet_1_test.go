package echo

import (
	"fmt"
	"testing"
)

func TestTrustPrivateNet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var v bool
	c := &ipChecker{}
	TrustPrivateNet(v)(c)
	if c.trustPrivateNet != v {
		t.Errorf("TrustPrivateNet(%v) = %v, want %v", v, c.trustPrivateNet, v)
	}
}
