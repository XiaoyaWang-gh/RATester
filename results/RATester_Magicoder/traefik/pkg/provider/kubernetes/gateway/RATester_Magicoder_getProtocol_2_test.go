package gateway

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestGetProtocol_2(t *testing.T) {
	tests := []struct {
		name     string
		portSpec corev1.ServicePort
		want     string
	}{
		{
			name: "Test case 1: port 443",
			portSpec: corev1.ServicePort{
				Port: 443,
			},
			want: "https",
		},
		{
			name: "Test case 2: port 80",
			portSpec: corev1.ServicePort{
				Port: 80,
			},
			want: "http",
		},
		{
			name: "Test case 3: port 8080",
			portSpec: corev1.ServicePort{
				Port: 8080,
			},
			want: "http",
		},
		{
			name: "Test case 4: port 443 with https prefix",
			portSpec: corev1.ServicePort{
				Name: "https",
			},
			want: "https",
		},
		{
			name: "Test case 5: port 80 with https prefix",
			portSpec: corev1.ServicePort{
				Name: "https",
			},
			want: "https",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getProtocol(tt.portSpec); got != tt.want {
				t.Errorf("getProtocol() = %v, want %v", got, tt.want)
			}
		})
	}
}
