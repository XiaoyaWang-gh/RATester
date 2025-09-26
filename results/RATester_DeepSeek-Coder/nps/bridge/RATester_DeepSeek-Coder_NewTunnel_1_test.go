package bridge

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestNewTunnel_1(t *testing.T) {
	type args struct {
		tunnelPort     int
		tunnelType     string
		ipVerify       bool
		runList        sync.Map
		disconnectTime int
	}
	tests := []struct {
		name string
		args args
		want *Bridge
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

			if got := NewTunnel(tt.args.tunnelPort, tt.args.tunnelType, tt.args.ipVerify, tt.args.runList, tt.args.disconnectTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}
