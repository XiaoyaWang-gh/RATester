package gateway

import (
	"fmt"
	"reflect"
	"testing"

	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestAllowedNamespaces_1(t *testing.T) {
	type args struct {
		gatewayNamespace string
		routeNamespaces  *gatev1.RouteNamespaces
	}
	tests := []struct {
		name    string
		args    args
		want    []string
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

			p := &Provider{}
			got, err := p.allowedNamespaces(tt.args.gatewayNamespace, tt.args.routeNamespaces)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.allowedNamespaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.allowedNamespaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
