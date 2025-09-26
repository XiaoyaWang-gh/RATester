package conn

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSetWriteDeadline_1(t *testing.T) {
	type fields struct {
		Conn net.Conn
		Rb   []byte
	}
	type args struct {
		deadline time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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
			if err := s.SetWriteDeadline(tt.args.deadline); (err != nil) != tt.wantErr {
				t.Errorf("Conn.SetWriteDeadline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
