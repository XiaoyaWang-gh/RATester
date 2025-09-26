package proxy

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetWorkConnFromPool_1(t *testing.T) {
	type args struct {
		src net.Addr
		dst net.Addr
	}
	tests := []struct {
		name         string
		args         args
		wantWorkConn net.Conn
		wantErr      bool
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

			pxy := &BaseProxy{}
			gotWorkConn, err := pxy.GetWorkConnFromPool(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseProxy.GetWorkConnFromPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotWorkConn, tt.wantWorkConn) {
				t.Errorf("BaseProxy.GetWorkConnFromPool() = %v, want %v", gotWorkConn, tt.wantWorkConn)
			}
		})
	}
}
