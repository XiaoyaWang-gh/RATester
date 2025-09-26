package gateway

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestCreateRequestRedirect_1(t *testing.T) {
	type args struct {
		filter    *gatev1.HTTPRequestRedirectFilter
		pathMatch gatev1.HTTPPathMatch
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

			if got := createRequestRedirect(tt.args.filter, tt.args.pathMatch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createRequestRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}
