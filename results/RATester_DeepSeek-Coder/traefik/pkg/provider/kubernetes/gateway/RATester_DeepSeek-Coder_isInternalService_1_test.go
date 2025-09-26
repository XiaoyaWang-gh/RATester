package gateway

import (
	"fmt"
	"testing"

	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestIsInternalService_1(t *testing.T) {
	type args struct {
		ref gatev1.BackendRef
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

			if got := isInternalService(tt.args.ref); got != tt.want {
				t.Errorf("isInternalService() = %v, want %v", got, tt.want)
			}
		})
	}
}
