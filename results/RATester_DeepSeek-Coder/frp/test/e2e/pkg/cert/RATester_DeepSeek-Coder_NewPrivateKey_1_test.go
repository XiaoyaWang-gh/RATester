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

	t.Parallel()

	priv, err := NewPrivateKey()
	if err != nil {
		t.Fatalf("NewPrivateKey() failed: %v", err)
	}

	if priv == nil {
		t.Fatal("NewPrivateKey() returned nil")
	}

	if priv.D == nil {
		t.Errorf("NewPrivateKey() returned a private key with nil D")
	}

	if priv.Primes == nil {
		t.Errorf("NewPrivateKey() returned a private key with nil Primes")
	}

	if len(priv.Primes) < 2 {
		t.Errorf("NewPrivateKey() returned a private key with less than 2 primes")
	}

	if priv.N == nil {
		t.Errorf("NewPrivateKey() returned a private key with nil N")
	}

	if priv.E != 0 {
		t.Errorf("NewPrivateKey() returned a private key with non-zero E")
	}
}
