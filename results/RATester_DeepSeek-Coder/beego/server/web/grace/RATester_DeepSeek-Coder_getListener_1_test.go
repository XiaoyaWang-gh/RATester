package grace

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetListener_1(t *testing.T) {
	type args struct {
		laddr string
	}
	tests := []struct {
		name    string
		srv     *Server
		args    args
		wantL   net.Listener
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

			gotL, err := tt.srv.getListener(tt.args.laddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.getListener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("Server.getListener() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
