package common

import (
	"fmt"
	"testing"
)

func TestTestUdpPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if !TestUdpPort(1234) {
		t.Errorf("TestUdpPort(1234) = false, want true")
	}
}
