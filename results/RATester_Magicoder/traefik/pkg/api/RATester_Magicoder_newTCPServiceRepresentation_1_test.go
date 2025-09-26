package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestnewTCPServiceRepresentation_1(t *testing.T) {
	type args struct {
		name string
		si   *runtime.TCPServiceInfo
	}
	tests := []struct {
		name string
		args args
		want tcpServiceRepresentation
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

			if got := newTCPServiceRepresentation(tt.args.name, tt.args.si); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTCPServiceRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}
