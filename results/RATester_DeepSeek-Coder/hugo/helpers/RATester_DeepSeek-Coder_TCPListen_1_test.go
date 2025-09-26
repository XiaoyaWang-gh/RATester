package helpers

import (
	"fmt"
	"testing"
)

func TestTCPListen_1(t *testing.T) {
	t.Run("should return a valid listener and tcp address", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		l, addr, err := TCPListen()
		if err != nil {
			t.Fatalf("TCPListen() error = %v", err)
		}
		defer l.Close()

		if l == nil {
			t.Errorf("TCPListen() got = %v, want not nil", l)
		}
		if addr == nil {
			t.Errorf("TCPListen() got = %v, want not nil", addr)
		}
		if addr.IP == nil {
			t.Errorf("TCPListen() got = %v, want not nil", addr.IP)
		}
		if addr.Port == 0 {
			t.Errorf("TCPListen() got = %v, want not 0", addr.Port)
		}
	})
}
