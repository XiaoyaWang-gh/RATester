package fiber

import (
	"fmt"
	"os"
	"testing"
)

func TestIsChild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Setenv(envPreforkChildKey, "")
	if IsChild() {
		t.Errorf("Expected IsChild() to return false when environment variable is not set")
	}

	os.Setenv(envPreforkChildKey, "1")
	if !IsChild() {
		t.Errorf("Expected IsChild() to return true when environment variable is set to 1")
	}

	os.Setenv(envPreforkChildKey, "0")
	if IsChild() {
		t.Errorf("Expected IsChild() to return false when environment variable is set to 0")
	}

	os.Setenv(envPreforkChildKey, "other")
	if IsChild() {
		t.Errorf("Expected IsChild() to return false when environment variable is set to 'other'")
	}
}
