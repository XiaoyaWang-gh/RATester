package group

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewTCPGroupListener_1(t *testing.T) {
	type args struct {
		name  string
		group *TCPGroup
		addr  net.Addr
	}
	tests := []struct {
		name string
		args args
		want *TCPGroupListener
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

			if got := newTCPGroupListener(tt.args.name, tt.args.group, tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTCPGroupListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
