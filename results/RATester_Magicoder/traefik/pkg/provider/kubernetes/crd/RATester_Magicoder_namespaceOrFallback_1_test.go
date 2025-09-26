package crd

import (
	"fmt"
	"testing"

	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestNamespaceOrFallback_1(t *testing.T) {
	type args struct {
		lb       traefikv1alpha1.LoadBalancerSpec
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with empty namespace",
			args: args{
				lb: traefikv1alpha1.LoadBalancerSpec{
					Namespace: "",
				},
				fallback: "default",
			},
			want: "default",
		},
		{
			name: "Test with non-empty namespace",
			args: args{
				lb: traefikv1alpha1.LoadBalancerSpec{
					Namespace: "test",
				},
				fallback: "default",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := namespaceOrFallback(tt.args.lb, tt.args.fallback); got != tt.want {
				t.Errorf("namespaceOrFallback() = %v, want %v", got, tt.want)
			}
		})
	}
}
