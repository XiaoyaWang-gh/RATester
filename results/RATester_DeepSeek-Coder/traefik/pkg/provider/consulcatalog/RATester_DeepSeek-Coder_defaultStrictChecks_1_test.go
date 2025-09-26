package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestDefaultStrictChecks_1(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test case 1",
			want: []string{api.HealthPassing, api.HealthWarning},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := defaultStrictChecks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultStrictChecks() = %v, want %v", got, tt.want)
			}
		})
	}
}
