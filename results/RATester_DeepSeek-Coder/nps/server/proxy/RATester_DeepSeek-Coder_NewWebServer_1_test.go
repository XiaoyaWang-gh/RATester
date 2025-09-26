package proxy

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/bridge"
)

func TestNewWebServer_1(t *testing.T) {
	type args struct {
		bridge *bridge.Bridge
	}
	tests := []struct {
		name string
		args args
		want *WebServer
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

			if got := NewWebServer(tt.args.bridge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWebServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
