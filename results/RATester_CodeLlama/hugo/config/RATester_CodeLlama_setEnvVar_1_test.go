package config

import (
	"fmt"
	"testing"
)

func TestSetEnvVar_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vars := []string{"a=1", "b=2"}
	key := "c"
	value := "3"
	setEnvVar(&vars, key, value)
	if len(vars) != 3 {
		t.Errorf("Expected 3, got %d", len(vars))
	}
	if vars[0] != "a=1" {
		t.Errorf("Expected a=1, got %s", vars[0])
	}
	if vars[1] != "b=2" {
		t.Errorf("Expected b=2, got %s", vars[1])
	}
	if vars[2] != "c=3" {
		t.Errorf("Expected c=3, got %s", vars[2])
	}
}
