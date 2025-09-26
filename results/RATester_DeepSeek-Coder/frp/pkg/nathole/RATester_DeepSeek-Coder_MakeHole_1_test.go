package nathole

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestMakeHole_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		listenConn *net.UDPConn
		m          *msg.NatHoleResp
		key        []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *net.UDPConn
		want1   *net.UDPAddr
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

			got, got1, err := MakeHole(tt.args.ctx, tt.args.listenConn, tt.args.m, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeHole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeHole() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MakeHole() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
