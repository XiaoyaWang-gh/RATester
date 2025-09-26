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
			name: "Context_0",
			args: args{
				ingressClass:           "test",
				ingressClassAnnotation: "test",
			},
			want: true,
		},
		{
			name: "Context_1",
			args: args{
				ingressClass:           "",
				ingressClassAnnotation: "test",
			},
			want: false,
		},
		{
			name: "Context_2",
			args: args{
				ingressClass:           "test",
				ingressClassAnnotation: "",
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
