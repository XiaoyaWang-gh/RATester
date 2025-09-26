package ingress

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	netv1 "k8s.io/api/networking/v1"
)

func TestLoadRouter_1(t *testing.T) {
	type args struct {
		rule        netv1.IngressRule
		pa          netv1.HTTPIngressPath
		rtConfig    *RouterConfig
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want *dynamic.Router
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

			if got := loadRouter(tt.args.rule, tt.args.pa, tt.args.rtConfig, tt.args.serviceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
