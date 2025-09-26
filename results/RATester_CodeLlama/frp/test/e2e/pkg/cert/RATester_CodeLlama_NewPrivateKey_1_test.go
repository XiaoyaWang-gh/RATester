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

	// Arrange
	// Act
	priv, err := NewPrivateKey()
	// Assert
	if err != nil {
		t.Errorf("NewPrivateKey() error = %v", err)
		return
	}
	if priv == nil {
		t.Errorf("NewPrivateKey() priv = nil")
		return
	}
	if priv.D == nil {
		t.Errorf("NewPrivateKey() priv.D = nil")
		return
	}
	if priv.Primes == nil {
		t.Errorf("NewPrivateKey() priv.Primes = nil")
		return
	}
	if priv.Precomputed.Dp == nil {
		t.Errorf("NewPrivateKey() priv.Precomputed.Dp = nil")
		return
	}
	if priv.Precomputed.Dq == nil {
		t.Errorf("NewPrivateKey() priv.Precomputed.Dq = nil")
		return
	}
	if priv.Precomputed.Qinv == nil {
		t.Errorf("NewPrivateKey() priv.Precomputed.Qinv = nil")
		return
	}
}
