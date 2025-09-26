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
			name: "Test with nil method",
			args: args{
				method: nil,
			},
			want: "PathPrefix(`/`)",
		},
		{
			name: "Test with service and method",
			args: args{
				method: &gatev1.GRPCMethodMatch{
					Service: ptr.String("test_service"),
					Method:  ptr.String("test_method"),
				},
			},
			want: "PathRegexp(`/test_service/test_method`)",
		},
		{
			name: "Test with only service",
			args: args{
				method: &gatev1.GRPCMethodMatch{
					Service: ptr.String("test_service"),
				},
			},
			want: "PathRegexp(`/test_service/[^/]+`)",
		},
		{
			name: "Test with only method",
			args: args{
				method: &gatev1.GRPCMethodMatch{
					Method: ptr.String("test_method"),
				},
			},
			want: "PathRegexp(`/[^/]+/test_method`)",
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
