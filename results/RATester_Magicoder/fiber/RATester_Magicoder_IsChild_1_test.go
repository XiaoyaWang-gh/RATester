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

	os.Setenv(envPreforkChildKey, "1")
	if !IsChild() {
		t.Errorf("Expected IsChild() to return true, but it returned false")
	}
	os.Unsetenv(envPreforkChildKey)
	if IsChild() {
		t.Errorf("Expected IsChild() to return false, but it returned true")
	}
}
