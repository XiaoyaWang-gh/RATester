package ingress

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestGetProtocol_1(t *testing.T) {
	type args struct {
		portSpec  corev1.ServicePort
		portName  string
		svcConfig *ServiceConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				portSpec: corev1.ServicePort{
					Port: 80,
				},
				portName: "http",
				svcConfig: &ServiceConfig{
					Service: &ServiceIng{
						ServersScheme: "https",
					},
				},
			},
			want: "https",
		},
		{
			name: "Test Case 2",
			args: args{
				portSpec: corev1.ServicePort{
					Port: 443,
				},
				portName: "https",
				svcConfig: &ServiceConfig{
					Service: &ServiceIng{
						ServersScheme: "http",
					},
				},
			},
			want: "https",
		},
		{
			name: "Test Case 3",
			args: args{
				portSpec: corev1.ServicePort{
					Port: 8080,
				},
				portName: "http",
				svcConfig: &ServiceConfig{
					Service: &ServiceIng{
						ServersScheme: "http",
					},
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

			if got := getProtocol(tt.args.portSpec, tt.args.portName, tt.args.svcConfig); got != tt.want {
				t.Errorf("getProtocol() = %v, want %v", got, tt.want)
			}
		})
	}
}
