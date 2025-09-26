package conn

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestLocalAddr_1(t *testing.T) {
	type fields struct {
		Conn net.Conn
		Rb   []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   net.Addr
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
			if got := s.LocalAddr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.LocalAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
