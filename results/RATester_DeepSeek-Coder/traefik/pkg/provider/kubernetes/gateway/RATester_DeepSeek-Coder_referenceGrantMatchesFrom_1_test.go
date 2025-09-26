package gateway

import (
	"fmt"
	"testing"

	gatev1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestReferenceGrantMatchesFrom_1(t *testing.T) {
	type args struct {
		referenceGrant *gatev1beta1.ReferenceGrant
		group          string
		kind           string
		namespace      string
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

			if got := referenceGrantMatchesFrom(tt.args.referenceGrant, tt.args.group, tt.args.kind, tt.args.namespace); got != tt.want {
				t.Errorf("referenceGrantMatchesFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
