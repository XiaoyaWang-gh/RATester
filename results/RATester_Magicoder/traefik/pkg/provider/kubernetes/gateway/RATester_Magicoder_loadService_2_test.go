package gateway

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestLoadService_2(t *testing.T) {
	type args struct {
		route      *gatev1.HTTPRoute
		backendRef gatev1.HTTPBackendRef
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 *dynamic.Service
		want2 *metav1.Condition
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
			got, got1, got2 := p.loadService(tt.args.route, tt.args.backendRef)
			if got != tt.want {
				t.Errorf("loadService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("loadService() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("loadService() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
