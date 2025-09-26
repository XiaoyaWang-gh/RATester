package common

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnvMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Setenv("TEST_KEY", "TEST_VALUE")
	defer os.Unsetenv("TEST_KEY")

	envMap := GetEnvMap()

	if val, ok := envMap["TEST_KEY"]; !ok || val != "TEST_VALUE" {
		t.Errorf("Expected 'TEST_KEY' to be 'TEST_VALUE', got %s", val)
	}
}
