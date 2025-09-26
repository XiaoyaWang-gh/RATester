package acme

import (
	"fmt"
	"testing"
)

func TestPresent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	challenge := &ChallengeTLSALPN{
		chans: make(map[string]chan struct{}),
		certs: make(map[string]*Certificate),
	}

	domain := "example.com"
	keyAuth := "keyAuth"

	err := challenge.Present(domain, "", keyAuth)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if _, ok := challenge.certs[keyAuth]; !ok {
		t.Errorf("Expected certificate to be stored, but it was not")
	}

	if _, ok := challenge.chans[string(challenge.certs[keyAuth].Certificate)]; !ok {
		t.Errorf("Expected channel to be stored, but it was not")
	}
}
