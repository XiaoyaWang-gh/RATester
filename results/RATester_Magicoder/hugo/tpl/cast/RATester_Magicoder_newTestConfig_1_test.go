package cast

import (
	"fmt"
	"testing"
)

func TestnewTestConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := newTestConfig()

	// Test case 1: Check if the contentDir is set correctly
	expectedContentDir := "content"
	actualContentDir := cfg.GetString("contentDir")
	if actualContentDir != expectedContentDir {
		t.Errorf("Expected contentDir to be %s, but got %s", expectedContentDir, actualContentDir)
	}

	// Test case 2: Check if the default value is returned for a non-existent key
	expectedDefaultValue := ""
	actualDefaultValue := cfg.GetString("nonExistentKey")
	if actualDefaultValue != expectedDefaultValue {
		t.Errorf("Expected default value to be %s, but got %s", expectedDefaultValue, actualDefaultValue)
	}

	// Test case 3: Check if the IsSet method correctly identifies if a key exists
	expectedIsSet := true
	actualIsSet := cfg.IsSet("contentDir")
	if actualIsSet != expectedIsSet {
		t.Errorf("Expected IsSet to return %t, but got %t", expectedIsSet, actualIsSet)
	}

	// Test case 4: Check if the IsSet method correctly identifies if a non-existent key exists
	expectedIsSet = false
	actualIsSet = cfg.IsSet("nonExistentKey")
	if actualIsSet != expectedIsSet {
		t.Errorf("Expected IsSet to return %t, but got %t", expectedIsSet, actualIsSet)
	}
}
