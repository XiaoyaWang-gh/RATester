package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestGetAddStatus_1(t *testing.T) {
	type fields struct {
		Conn net.Conn
		Rb   []byte
	}
	tests := []struct {
		name   string
		fields fields
		wantB  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &Conn{
				Conn: tt.fields.Conn,
				Rb:   tt.fields.Rb,
			}
			if gotB := s.GetAddStatus(); gotB != tt.wantB {
				t.Errorf("Conn.GetAddStatus() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
