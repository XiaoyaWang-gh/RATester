package gin

import (
	"fmt"
	"testing"
)

func TestSetTrustedProxies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	if len(engine.trustedProxies) != 1 {
		t.Errorf("Expected 1 trusted proxy, got %d", len(engine.trustedProxies))
	}
	if engine.trustedProxies[0] != "127.0.0.1" {
		t.Errorf("Expected trusted proxy to be 127.0.0.1, got %s", engine.trustedProxies[0])
	}
}
