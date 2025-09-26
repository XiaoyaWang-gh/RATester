package gin

import (
	"fmt"
	"testing"
)

func TestParseTrustedProxies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		trustedProxies: []string{"192.0.2.0/29", "203.0.113.0/24"},
	}

	err := engine.parseTrustedProxies()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(engine.trustedCIDRs) != 2 {
		t.Errorf("Expected 2 trusted CIDRs, got %d", len(engine.trustedCIDRs))
	}

	expectedCIDRs := []string{"192.0.2.0/29", "203.0.113.0/24"}
	for i, cidr := range engine.trustedCIDRs {
		if cidr.String() != expectedCIDRs[i] {
			t.Errorf("Expected trusted CIDR %s, got %s", expectedCIDRs[i], cidr.String())
		}
	}
}
