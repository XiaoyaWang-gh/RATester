package acme

import (
	"fmt"
	"testing"
)

func TestNewChallengeHTTP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	challengeHTTP := NewChallengeHTTP()
	if challengeHTTP == nil {
		t.Errorf("NewChallengeHTTP() = nil")
	}
}
