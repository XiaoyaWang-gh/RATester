package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestNewTcpListenerAndProcess_1(t *testing.T) {
	type args struct {
		addr string
		f    func(c net.Conn)
		l    *net.Listener
	}
	tests := []struct {
		name    string
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

			if err := NewTcpListenerAndProcess(tt.args.addr, tt.args.f, tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("NewTcpListenerAndProcess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
