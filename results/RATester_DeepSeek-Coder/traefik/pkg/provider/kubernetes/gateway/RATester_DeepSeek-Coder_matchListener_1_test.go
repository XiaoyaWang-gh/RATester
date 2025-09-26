package gateway

import (
	"fmt"
	"testing"

	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestMatchListener_1(t *testing.T) {
	type args struct {
		listener       gatewayListener
		routeNamespace string
		parentRef      gatev1.ParentReference
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := matchListener(tt.args.listener, tt.args.routeNamespace, tt.args.parentRef); got != tt.want {
				t.Errorf("matchListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
