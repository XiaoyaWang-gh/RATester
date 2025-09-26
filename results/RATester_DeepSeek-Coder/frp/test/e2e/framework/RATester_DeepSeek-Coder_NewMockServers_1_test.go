package framework

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/port"
)

func TestNewMockServers_1(t *testing.T) {
	type args struct {
		portAllocator *port.Allocator
	}
	tests := []struct {
		name string
		args args
		want *MockServers
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

			if got := NewMockServers(tt.args.portAllocator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
