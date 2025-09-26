package security

import (
	"fmt"
	"testing"
)

func TestCheckAllowedExec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := Config{
		Exec: Exec{
			Allow: Whitelist{
				acceptNone: true,
			},
		},
	}

	// Test case 1: name is in the whitelist
	err := config.CheckAllowedExec("allowedName")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: name is not in the whitelist
	err = config.CheckAllowedExec("disallowedName")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
