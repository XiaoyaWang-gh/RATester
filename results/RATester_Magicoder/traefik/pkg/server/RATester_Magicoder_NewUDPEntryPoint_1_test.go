package server

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewUDPEntryPoint_1(t *testing.T) {
	type args struct {
		cfg *static.EntryPoint
	}
	tests := []struct {
		name    string
		args    args
		want    *UDPEntryPoint
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

			got, err := NewUDPEntryPoint(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUDPEntryPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUDPEntryPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
