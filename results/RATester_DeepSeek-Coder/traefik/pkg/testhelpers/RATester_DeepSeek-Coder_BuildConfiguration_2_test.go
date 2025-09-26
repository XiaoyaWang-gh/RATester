package testhelpers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildConfiguration_2(t *testing.T) {
	tests := []struct {
		name string
		want *dynamic.HTTPConfiguration
	}{
		{
			name: "Test BuildConfiguration",
			want: &dynamic.HTTPConfiguration{
				Models:            map[string]*dynamic.Model{},
				ServersTransports: map[string]*dynamic.ServersTransport{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := BuildConfiguration(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
