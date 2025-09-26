package acme

import (
	"fmt"
	"testing"
)

func TestCleanUp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &ChallengeTLSALPN{
		certs: make(map[string]*Certificate),
	}

	keyAuth := "keyAuth"
	provider.certs[keyAuth] = &Certificate{}

	err := provider.CleanUp("example.com", "", keyAuth)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if _, ok := provider.certs[keyAuth]; ok {
		t.Errorf("Expected certificate to be deleted, but it still exists")
	}
}
