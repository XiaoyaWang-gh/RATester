package server

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/bridge"
	"ehang.io/nps/lib/file"
	"ehang.io/nps/server/proxy"
)

func TestNewMode_1(t *testing.T) {
	type args struct {
		Bridge *bridge.Bridge
		c      *file.Tunnel
	}
	tests := []struct {
		name string
		args args
		want proxy.Service
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

			if got := NewMode(tt.args.Bridge, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
