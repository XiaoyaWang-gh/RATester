package provider

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddService_1(t *testing.T) {
	type args struct {
		configuration *dynamic.HTTPConfiguration
		serviceName   string
		service       *dynamic.Service
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := AddService(tt.args.configuration, tt.args.serviceName, tt.args.service); got != tt.want {
				t.Errorf("AddService() = %v, want %v", got, tt.want)
			}
		})
	}
}
