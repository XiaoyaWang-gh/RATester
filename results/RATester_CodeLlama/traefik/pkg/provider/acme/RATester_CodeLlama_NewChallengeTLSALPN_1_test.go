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

	challenge := NewChallengeTLSALPN()
	if challenge == nil {
		t.Errorf("NewChallengeTLSALPN() = %v, want %v", challenge, &ChallengeTLSALPN{})
	}
}
