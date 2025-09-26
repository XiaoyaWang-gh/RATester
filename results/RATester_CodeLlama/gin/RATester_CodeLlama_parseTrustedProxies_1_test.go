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

	engine := &Engine{}
	err := engine.parseTrustedProxies()
	if err != nil {
		t.Errorf("parseTrustedProxies() error = %v", err)
		return
	}
}
