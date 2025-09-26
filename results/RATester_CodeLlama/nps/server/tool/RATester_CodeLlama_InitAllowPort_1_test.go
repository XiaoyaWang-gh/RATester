package tool

import (
	"fmt"
	"testing"
)

func TestInitAllowPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	InitAllowPort()
	if len(ports) != 1 {
		t.Errorf("InitAllowPort() = %v, want %v", len(ports), 1)
	}
}
