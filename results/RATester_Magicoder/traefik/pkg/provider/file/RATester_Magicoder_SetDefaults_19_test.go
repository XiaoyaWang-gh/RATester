package file

import (
	"fmt"
	"testing"
)

func TestSetDefaults_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{}
	provider.SetDefaults()

	if !provider.Watch {
		t.Error("Expected Watch to be true, but it was false")
	}

	if provider.Filename != "" {
		t.Error("Expected Filename to be empty, but it was not")
	}
}
