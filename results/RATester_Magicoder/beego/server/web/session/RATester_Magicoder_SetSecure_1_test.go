package session

import (
	"fmt"
	"testing"
)

func TestSetSecure_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &Manager{
		config: &ManagerConfig{
			Secure: false,
		},
	}

	manager.SetSecure(true)

	if manager.config.Secure != true {
		t.Errorf("Expected Secure to be true, but got %v", manager.config.Secure)
	}
}
