package gin

import (
	"fmt"
	"testing"
)

func TestPrepareTrustedCIDRs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		trustedProxies: []string{"192.0.2.0/29", "203.0.113.0/24"},
	}

	cidr, err := engine.prepareTrustedCIDRs()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(cidr) != 2 {
		t.Errorf("Expected 2 CIDRs, got %d", len(cidr))
	}

	expectedCIDRs := []string{"192.0.2.0/29", "203.0.113.0/24"}
	for i, c := range cidr {
		if c.String() != expectedCIDRs[i] {
			t.Errorf("Expected CIDR %s, got %s", expectedCIDRs[i], c.String())
		}
	}
}
