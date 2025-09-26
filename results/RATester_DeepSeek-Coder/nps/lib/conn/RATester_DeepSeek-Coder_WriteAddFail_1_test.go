package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestWriteAddFail_1(t *testing.T) {
	type fields struct {
		Conn net.Conn
		Rb   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
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
			if err := s.WriteAddFail(); (err != nil) != tt.wantErr {
				t.Errorf("Conn.WriteAddFail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
