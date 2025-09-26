package pmux

import (
	"fmt"
	"net"
	"testing"
)

func Testprocess_2(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
		want *PortMux
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

			pMux := &PortMux{}
			pMux.process(tt.args.conn)
		})
	}
}
