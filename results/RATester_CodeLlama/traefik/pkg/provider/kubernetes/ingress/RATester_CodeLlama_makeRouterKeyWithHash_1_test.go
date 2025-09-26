package ingress

import (
	"fmt"
	"testing"
)

func TestMakeRouterKeyWithHash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "test"
	rule := "test"
	dupKey, err := makeRouterKeyWithHash(key, rule)
	if err != nil {
		t.Errorf("makeRouterKeyWithHash error: %v", err)
	}
	if dupKey != "test-00000000000000000000000000000000" {
		t.Errorf("makeRouterKeyWithHash error: %v", dupKey)
	}
}
