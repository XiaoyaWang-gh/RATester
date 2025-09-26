package fiber

import (
	"fmt"
	"testing"
)

func TesthandleTrustedProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}

	// Test case 1: Valid CIDR
	app.handleTrustedProxy("192.168.0.0/24")
	if len(app.config.trustedProxyRanges) != 1 {
		t.Errorf("Expected 1 trusted proxy range, got %d", len(app.config.trustedProxyRanges))
	}

	// Test case 2: Invalid CIDR
	app.handleTrustedProxy("192.168.0.0/33")
	if len(app.config.trustedProxyRanges) != 1 {
		t.Errorf("Expected 1 trusted proxy range, got %d", len(app.config.trustedProxyRanges))
	}

	// Test case 3: Valid IP
	app.handleTrustedProxy("192.168.0.1")
	if len(app.config.trustedProxiesMap) != 1 {
		t.Errorf("Expected 1 trusted proxy, got %d", len(app.config.trustedProxiesMap))
	}
}
