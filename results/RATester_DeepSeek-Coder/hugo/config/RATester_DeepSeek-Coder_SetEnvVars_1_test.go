package config

import (
	"fmt"
	"os"
	"testing"
)

func TestSetEnvVars_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	oldVars := []string{}
	newVars := []string{"KEY1", "VALUE1", "KEY2", "VALUE2"}

	SetEnvVars(&oldVars, newVars...)

	for i := 0; i < len(newVars); i += 2 {
		if os.Getenv(newVars[i]) != newVars[i+1] {
			t.Errorf("Expected %s to be %s", newVars[i], newVars[i+1])
		}
	}

	// Restore old environment variables
	for _, key := range oldVars {
		os.Unsetenv(key)
	}
}
