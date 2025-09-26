package config

import (
	"fmt"
	"testing"
)

func TestWithSSHMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	options := &registerFlagOptions{}
	WithSSHMode()(options)

	if !options.sshMode {
		t.Error("Expected sshMode to be true, but it was false")
	}
}
