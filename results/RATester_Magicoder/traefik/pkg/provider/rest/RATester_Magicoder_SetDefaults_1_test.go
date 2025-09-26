package rest

import (
	"fmt"
	"testing"
)

func TestSetDefaults_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Insecure: false,
	}

	provider.SetDefaults()

	if provider.Insecure != false {
		t.Errorf("Expected Insecure to be false, but got %v", provider.Insecure)
	}
}
