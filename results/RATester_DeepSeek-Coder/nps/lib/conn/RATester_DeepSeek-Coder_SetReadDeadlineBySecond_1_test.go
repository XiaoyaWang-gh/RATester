package conn

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSetReadDeadlineBySecond_1(t *testing.T) {
	type fields struct {
		Conn net.Conn
		Rb   []byte
	}
	type args struct {
		t time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			s.SetReadDeadlineBySecond(tt.args.t)
		})
	}
}
