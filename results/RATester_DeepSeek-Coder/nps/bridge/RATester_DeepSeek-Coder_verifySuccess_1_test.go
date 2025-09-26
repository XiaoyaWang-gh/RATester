package bridge

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestVerifySuccess_1(t *testing.T) {
	type args struct {
		c *conn.Conn
	}

	mockConn := &conn.Conn{
		Conn: nil,
		Rb:   []byte{},
	}

	tests := []struct {
		name string
		s    *Bridge
		args args
	}{
		{
			name: "Test verifySuccess",
			s:    &Bridge{},
			args: args{c: mockConn},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.s.verifySuccess(tt.args.c)
		})
	}
}
