package gateway

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestGetProtocol_2(t *testing.T) {
	type args struct {
		portSpec corev1.ServicePort
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				portSpec: corev1.ServicePort{
					Port: 443,
				},
			},
			want: "https",
		},
		{
			name: "test case 2",
			args: args{
				portSpec: corev1.ServicePort{
					Name: "https",
				},
			},
			want: "https",
		},
		{
			name: "test case 3",
			args: args{
				portSpec: corev1.ServicePort{
					Port: 80,
				},
			},
			want: "http",
		},
		{
			name: "test case 4",
			args: args{
				portSpec: corev1.ServicePort{
					Name: "http",
				},
			},
			want: "http",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getProtocol(tt.args.portSpec); got != tt.want {
				t.Errorf("getProtocol() = %v, want %v", got, tt.want)
			}
		})
	}
}
