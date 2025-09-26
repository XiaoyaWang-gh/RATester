package gin

import (
	"fmt"
	"testing"
)

func TestSecureJsonPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{}
	prefix := "test_prefix"
	engine.SecureJsonPrefix(prefix)

	if engine.secureJSONPrefix != prefix {
		t.Errorf("Expected secureJSONPrefix to be %s, but got %s", prefix, engine.secureJSONPrefix)
	}
}
