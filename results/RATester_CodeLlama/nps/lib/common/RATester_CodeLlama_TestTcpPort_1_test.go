package common

import (
	"fmt"
	"testing"
)

func TestTestTcpPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if !TestTcpPort(8080) {
		t.Errorf("TestTcpPort(8080) = false, want true")
	}
}
