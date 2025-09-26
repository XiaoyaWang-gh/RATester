package web

import (
	"fmt"
	"testing"
)

func TestCheckXSRFCookie_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		EnableXSRF: false,
	}

	result := ctrl.CheckXSRFCookie()

	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}
