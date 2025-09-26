package config

import (
	"fmt"
	"testing"
)

func TestsetEnvVar_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vars := []string{"FOO=bar", "BAZ=qux"}
	setEnvVar(&vars, "FOO", "baz")

	if vars[0] != "FOO=baz" {
		t.Errorf("Expected 'FOO=baz', got '%s'", vars[0])
	}

	setEnvVar(&vars, "NEW", "value")

	if vars[2] != "NEW=value" {
		t.Errorf("Expected 'NEW=value', got '%s'", vars[2])
	}
}
