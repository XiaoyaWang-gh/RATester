package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestNewUDPServiceRepresentation_1(t *testing.T) {
	type args struct {
		name string
		si   *runtime.UDPServiceInfo
	}
	tests := []struct {
		name string
		args args
		want udpServiceRepresentation
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

			if got := newUDPServiceRepresentation(tt.args.name, tt.args.si); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newUDPServiceRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}
