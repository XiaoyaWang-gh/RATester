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

	o := &registerFlagOptions{}
	WithSSHMode()(o)
	if o.sshMode != true {
		t.Errorf("WithSSHMode() failed")
	}
}
