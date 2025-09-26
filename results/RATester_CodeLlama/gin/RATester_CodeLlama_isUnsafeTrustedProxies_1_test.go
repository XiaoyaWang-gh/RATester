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

	engine := &Engine{}
	engine.trustedProxies = []string{"0.0.0.0", "::"}
	if !engine.isUnsafeTrustedProxies() {
		t.Errorf("engine.isUnsafeTrustedProxies() = %v, want %v", engine.isUnsafeTrustedProxies(), true)
	}
}
