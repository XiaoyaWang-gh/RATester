package gateway

import (
	"fmt"
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestAllowedRouteKinds_1(t *testing.T) {
	type args struct {
		gateway        *gatev1.Gateway
		listener       gatev1.Listener
		supportedKinds []gatev1.RouteGroupKind
	}
	tests := []struct {
		name  string
		args  args
		want  []gatev1.RouteGroupKind
		want1 []metav1.Condition
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

			got, got1 := allowedRouteKinds(tt.args.gateway, tt.args.listener, tt.args.supportedKinds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allowedRouteKinds() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("allowedRouteKinds() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
