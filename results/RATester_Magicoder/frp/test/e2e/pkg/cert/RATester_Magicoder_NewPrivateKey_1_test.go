package cert

import (
	"fmt"
	"testing"
)

func TestNewPrivateKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key, err := NewPrivateKey()
	if err != nil {
		t.Errorf("NewPrivateKey() failed: %v", err)
	}

	if key == nil {
		t.Error("NewPrivateKey() returned nil")
	}

	if key.D == nil {
		t.Error("NewPrivateKey() returned key with nil D")
	}

	if len(key.Primes) < 2 {
		t.Errorf("NewPrivateKey() returned key with less than 2 primes: %d", len(key.Primes))
	}
}
