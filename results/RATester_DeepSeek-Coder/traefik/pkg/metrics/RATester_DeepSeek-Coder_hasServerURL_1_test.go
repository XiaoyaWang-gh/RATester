package metrics

import (
	"fmt"
	"testing"
)

func TestHasServerURL_1(t *testing.T) {
	type args struct {
		serviceName string
		serverURL   string
	}
	tests := []struct {
		name string
		d    *dynamicConfig
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			d: &dynamicConfig{
				services: map[string]map[string]bool{
					"service1": {"http://localhost:8080": true},
				},
			},
			args: args{
				serviceName: "service1",
				serverURL:   "http://localhost:8080",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			d: &dynamicConfig{
				services: map[string]map[string]bool{
					"service1": {"http://localhost:8080": true},
				},
			},
			args: args{
				serviceName: "service1",
				serverURL:   "http://localhost:8081",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			d: &dynamicConfig{
				services: map[string]map[string]bool{
					"service1": {"http://localhost:8080": true},
				},
			},
			args: args{
				serviceName: "service2",
				serverURL:   "http://localhost:8080",
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

			if got := tt.d.hasServerURL(tt.args.serviceName, tt.args.serverURL); got != tt.want {
				t.Errorf("dynamicConfig.hasServerURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
