package gin

import (
	"fmt"
	"testing"
)

func TestIsUnsafeTrustedProxies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		trustedProxies: []string{"0.0.0.0", "::"},
	}

	result := engine.isUnsafeTrustedProxies()
	if !result {
		t.Errorf("Expected true, got false")
	}

	engine.trustedProxies = []string{"127.0.0.1", "::1"}
	result = engine.isUnsafeTrustedProxies()
	if result {
		t.Errorf("Expected false, got true")
	}
}
