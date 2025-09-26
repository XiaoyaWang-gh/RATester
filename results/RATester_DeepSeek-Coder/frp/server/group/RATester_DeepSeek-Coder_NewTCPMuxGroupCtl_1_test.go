package group

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/util/tcpmux"
)

func TestNewTCPMuxGroupCtl_1(t *testing.T) {
	type args struct {
		tcpMuxHTTPConnectMuxer *tcpmux.HTTPConnectTCPMuxer
	}
	tests := []struct {
		name string
		args args
		want *TCPMuxGroupCtl
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

			if got := NewTCPMuxGroupCtl(tt.args.tcpMuxHTTPConnectMuxer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTCPMuxGroupCtl() = %v, want %v", got, tt.want)
			}
		})
	}
}
