package acme

import (
	"fmt"
	"testing"
)

func TestIsDomainAlreadyChecked_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	domainToCheck := "example.com"
	existentDomains := []string{"example.com,example.net"}
	if !isDomainAlreadyChecked(domainToCheck, existentDomains) {
		t.Errorf("isDomainAlreadyChecked() = %v, want %v", false, true)
	}
}
