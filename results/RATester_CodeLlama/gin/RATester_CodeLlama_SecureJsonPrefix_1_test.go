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

	engine := New()
	engine.SecureJsonPrefix("")
	if engine.secureJSONPrefix != "" {
		t.Errorf("SecureJsonPrefix() error, expected: '', actual: '%s'", engine.secureJSONPrefix)
	}

	engine.SecureJsonPrefix(" ")
	if engine.secureJSONPrefix != " " {
		t.Errorf("SecureJsonPrefix() error, expected: ' ', actual: '%s'", engine.secureJSONPrefix)
	}

	engine.SecureJsonPrefix("abc")
	if engine.secureJSONPrefix != "abc" {
		t.Errorf("SecureJsonPrefix() error, expected: 'abc', actual: '%s'", engine.secureJSONPrefix)
	}
}
