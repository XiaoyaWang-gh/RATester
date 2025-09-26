package provider

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddServiceUDP_1(t *testing.T) {
	type args struct {
		configuration *dynamic.UDPConfiguration
		serviceName   string
		service       *dynamic.UDPService
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

			if got := AddServiceUDP(tt.args.configuration, tt.args.serviceName, tt.args.service); got != tt.want {
				t.Errorf("AddServiceUDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
