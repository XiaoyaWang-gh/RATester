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
		t.Error("Expected a non-nil ChallengeHTTP, but got nil")
	}

	if challengeHTTP.httpChallenges == nil {
		t.Error("Expected a non-nil httpChallenges, but got nil")
	}
}
