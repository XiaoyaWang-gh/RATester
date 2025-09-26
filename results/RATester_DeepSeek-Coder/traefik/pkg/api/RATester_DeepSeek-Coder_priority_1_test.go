package api

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestPriority_1(t *testing.T) {
	type fields struct {
		UDPRouterInfo *runtime.UDPRouterInfo
		Name          string
		Provider      string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
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

			r := udpRouterRepresentation{
				UDPRouterInfo: tt.fields.UDPRouterInfo,
				Name:          tt.fields.Name,
				Provider:      tt.fields.Provider,
			}
			if got := r.priority(); got != tt.want {
				t.Errorf("udpRouterRepresentation.priority() = %v, want %v", got, tt.want)
			}
		})
	}
}
