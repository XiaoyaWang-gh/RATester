package acme

import (
	"fmt"
	"testing"
)

func TestNewChallengeTLSALPN_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	challengeTLSALPN := NewChallengeTLSALPN()

	if challengeTLSALPN == nil {
		t.Error("Expected a non-nil challengeTLSALPN, but got nil")
	}

	if challengeTLSALPN.chans == nil {
		t.Error("Expected a non-nil chans, but got nil")
	}

	if challengeTLSALPN.certs == nil {
		t.Error("Expected a non-nil certs, but got nil")
	}
}
