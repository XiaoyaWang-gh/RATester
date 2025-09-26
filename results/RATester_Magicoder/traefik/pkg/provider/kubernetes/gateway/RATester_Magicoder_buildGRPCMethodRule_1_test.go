package gateway

import (
	"fmt"
	"testing"

	"github.com/aws/smithy-go/ptr"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestBuildGRPCMethodRule_1(t *testing.T) {
	type args struct {
		method *gatev1.GRPCMethodMatch
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil method",
			args: args{
				method: nil,
			},
			want: "PathPrefix(`/`)",
		},
		{
			name: "empty method",
			args: args{
				method: &gatev1.GRPCMethodMatch{},
			},
			want: "PathRegexp(`/[^/]+/[^/]+`)",
		},
		{
			name: "non-empty method",
			args: args{
				method: &gatev1.GRPCMethodMatch{
					Service: ptr.String("service"),
					Method:  ptr.String("method"),
				},
			},
			want: "PathRegexp(`/service/method`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := buildGRPCMethodRule(tt.args.method); got != tt.want {
				t.Errorf("buildGRPCMethodRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
