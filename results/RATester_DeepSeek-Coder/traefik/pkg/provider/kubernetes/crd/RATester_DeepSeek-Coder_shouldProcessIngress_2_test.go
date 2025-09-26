package crd

import (
	"fmt"
	"testing"
)

func TestShouldProcessIngress_2(t *testing.T) {
	type args struct {
		ingressClass           string
		ingressClassAnnotation string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true when ingressClass is empty and ingressClassAnnotation is traefikDefaultIngressClass",
			args: args{
				ingressClass:           "",
				ingressClassAnnotation: traefikDefaultIngressClass,
			},
			want: true,
		},
		{
			name: "should return true when ingressClass equals ingressClassAnnotation",
			args: args{
				ingressClass:           "test",
				ingressClassAnnotation: "test",
			},
			want: true,
		},
		{
			name: "should return false when ingressClass is not empty and ingressClassAnnotation is not traefikDefaultIngressClass",
			args: args{
				ingressClass:           "test",
				ingressClassAnnotation: "notTraefik",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := shouldProcessIngress(tt.args.ingressClass, tt.args.ingressClassAnnotation); got != tt.want {
				t.Errorf("shouldProcessIngress() = %v, want %v", got, tt.want)
			}
		})
	}
}
