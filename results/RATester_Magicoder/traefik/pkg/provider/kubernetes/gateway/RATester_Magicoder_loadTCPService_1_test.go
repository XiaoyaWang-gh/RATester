package gateway

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
	gatev1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

func TestLoadTCPService_1(t *testing.T) {
	type args struct {
		route      *gatev1alpha2.TCPRoute
		backendRef gatev1.BackendRef
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 *dynamic.TCPService
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
			got, got1, got2 := p.loadTCPService(tt.args.route, tt.args.backendRef)
			if got != tt.want {
				t.Errorf("loadTCPService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("loadTCPService() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("loadTCPService() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
