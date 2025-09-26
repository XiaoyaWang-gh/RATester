package gin

import (
	"fmt"
	"testing"
)

func TestparseTrustedProxies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		trustedProxies: []string{"192.168.1.1", "192.168.1.2"},
	}

	err := engine.parseTrustedProxies()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(engine.trustedCIDRs) != 2 {
		t.Errorf("Expected 2 trusted CIDRs, got %d", len(engine.trustedCIDRs))
	}

	if engine.trustedCIDRs[0].String() != "192.168.1.1/32" {
		t.Errorf("Expected first trusted CIDR to be 192.168.1.1/32, got %s", engine.trustedCIDRs[0].String())
	}

	if engine.trustedCIDRs[1].String() != "192.168.1.2/32" {
		t.Errorf("Expected second trusted CIDR to be 192.168.1.2/32, got %s", engine.trustedCIDRs[1].String())
	}
}
