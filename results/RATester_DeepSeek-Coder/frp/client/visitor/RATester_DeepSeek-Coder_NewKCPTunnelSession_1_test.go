package visitor

import (
	"fmt"
	"testing"
)

func TestNewKCPTunnelSession_1(t *testing.T) {
	t.Run("NewKCPTunnelSession", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		session := NewKCPTunnelSession()
		if session == nil {
			t.Errorf("NewKCPTunnelSession() = %v, want not nil", session)
		}
	})
}
