package visitor

import (
	"fmt"
	"testing"
)

func TestNewKCPTunnelSession_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	session := NewKCPTunnelSession()
	if session == nil {
		t.Error("NewKCPTunnelSession() should not return nil")
	}
}
