package gateway

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestCreateResponseHeaderModifier_1(t *testing.T) {
	type args struct {
		filter *gatev1.HTTPHeaderFilter
	}
	tests := []struct {
		name string
		args args
		want *dynamic.Middleware
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

			if got := createResponseHeaderModifier(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createResponseHeaderModifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
