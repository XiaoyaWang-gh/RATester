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

	oldVars := os.Environ()
	keyValues := []string{"key1", "value1", "key2", "value2"}
	SetEnvVars(&oldVars, keyValues...)
	if len(os.Environ()) != len(oldVars)+len(keyValues)/2 {
		t.Errorf("SetEnvVars failed")
	}
}
