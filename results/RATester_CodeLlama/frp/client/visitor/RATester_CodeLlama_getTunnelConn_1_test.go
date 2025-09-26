package visitor

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetTunnelConn_1(t *testing.T) {
	type args struct {
		sv *XTCPVisitor
	}
	tests := []struct {
		name    string
		args    args
		want    net.Conn
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

			got, err := tt.args.sv.getTunnelConn()
			if (err != nil) != tt.wantErr {
				t.Errorf("getTunnelConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTunnelConn() got = %v, want %v", got, tt.want)
			}
		})
	}
}
