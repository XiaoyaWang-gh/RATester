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
		t.Fatal("Expected a non-nil ChallengeTLSALPN, got nil")
	}

	if challengeTLSALPN.chans == nil {
		t.Error("Expected chans to be initialized, got nil")
	}

	if challengeTLSALPN.certs == nil {
		t.Error("Expected certs to be initialized, got nil")
	}
}
