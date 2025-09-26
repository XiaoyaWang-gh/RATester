package gateway

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/provider/kubernetes/k8s"
)

func TestSetRouterTransform_2(t *testing.T) {
	type args struct {
		routerTransform k8s.RouterTransform
	}
	tests := []struct {
		name string
		p    *Provider
		args args
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

			tt.p.SetRouterTransform(tt.args.routerTransform)
		})
	}
}
