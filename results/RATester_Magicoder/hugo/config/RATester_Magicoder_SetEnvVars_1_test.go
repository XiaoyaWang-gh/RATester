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
	keyValues := []string{"TEST_KEY1", "TEST_VALUE1", "TEST_KEY2", "TEST_VALUE2"}

	SetEnvVars(&oldVars, keyValues...)

	// Test if the environment variables are set correctly
	value1 := os.Getenv("TEST_KEY1")
	if value1 != "TEST_VALUE1" {
		t.Errorf("Expected TEST_KEY1 to be TEST_VALUE1, but got %s", value1)
	}

	value2 := os.Getenv("TEST_KEY2")
	if value2 != "TEST_VALUE2" {
		t.Errorf("Expected TEST_KEY2 to be TEST_VALUE2, but got %s", value2)
	}

	// Test if the old environment variables are saved correctly
	if len(oldVars) != 0 {
		t.Errorf("Expected oldVars to be empty, but got %v", oldVars)
	}

	// Clean up the environment variables
	os.Unsetenv("TEST_KEY1")
	os.Unsetenv("TEST_KEY2")
}
