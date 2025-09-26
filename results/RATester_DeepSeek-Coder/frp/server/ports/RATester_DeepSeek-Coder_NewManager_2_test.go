package ports

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/config/types"
)

func TestNewManager_2(t *testing.T) {
	type args struct {
		netType    string
		bindAddr   string
		allowPorts []types.PortsRange
	}
	tests := []struct {
		name string
		args args
		want *Manager
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

			if got := NewManager(tt.args.netType, tt.args.bindAddr, tt.args.allowPorts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
