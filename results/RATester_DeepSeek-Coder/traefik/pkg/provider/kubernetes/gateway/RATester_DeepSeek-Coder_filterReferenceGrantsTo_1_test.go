package gateway

import (
	"fmt"
	"reflect"
	"testing"

	gatev1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestFilterReferenceGrantsTo_1(t *testing.T) {
	type args struct {
		referenceGrants []*gatev1beta1.ReferenceGrant
		group           string
		kind            string
		name            string
	}
	tests := []struct {
		name string
		args args
		want []*gatev1beta1.ReferenceGrant
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

			if got := filterReferenceGrantsTo(tt.args.referenceGrants, tt.args.group, tt.args.kind, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterReferenceGrantsTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
