package grace

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewServer_1(t *testing.T) {
	type args struct {
		addr    string
		handler http.Handler
		opts    []ServerOption
	}
	tests := []struct {
		name    string
		args    args
		wantSrv *Server
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

			gotSrv := NewServer(tt.args.addr, tt.args.handler, tt.args.opts...)
			if !reflect.DeepEqual(gotSrv, tt.wantSrv) {
				t.Errorf("NewServer() = %v, want %v", gotSrv, tt.wantSrv)
			}
		})
	}
}
