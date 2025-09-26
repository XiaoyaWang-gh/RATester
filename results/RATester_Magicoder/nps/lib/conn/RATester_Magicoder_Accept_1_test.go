package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestAccept_1(t *testing.T) {
	type args struct {
		l net.Listener
		f func(c net.Conn)
	}
	tests := []struct {
		name string
		args args
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

			Accept(tt.args.l, tt.args.f)
		})
	}
}
