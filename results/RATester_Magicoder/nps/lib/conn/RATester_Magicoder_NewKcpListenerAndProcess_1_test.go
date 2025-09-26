package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestNewKcpListenerAndProcess_1(t *testing.T) {
	type args struct {
		addr string
		f    func(c net.Conn)
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

			if err := NewKcpListenerAndProcess(tt.args.addr, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("NewKcpListenerAndProcess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
