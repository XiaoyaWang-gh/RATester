package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetTrustedProxies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{}

	trustedProxies := []string{"192.168.1.1", "192.168.1.2"}
	err := engine.SetTrustedProxies(trustedProxies)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(engine.trustedProxies, trustedProxies) {
		t.Errorf("Expected trustedProxies to be %v, but got %v", trustedProxies, engine.trustedProxies)
	}
}
