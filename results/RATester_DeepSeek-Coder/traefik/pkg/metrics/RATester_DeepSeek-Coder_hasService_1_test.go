package metrics

import (
	"fmt"
	"testing"
)

func TestHasService_1(t *testing.T) {
	type args struct {
		serviceName string
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
					"service1": {"url1": true, "url2": true},
					"service2": {"url3": true, "url4": true},
				},
			},
			args: args{serviceName: "service1"},
			want: true,
		},
		{
			name: "Test Case 2",
			d: &dynamicConfig{
				services: map[string]map[string]bool{
					"service1": {"url1": true, "url2": true},
					"service2": {"url3": true, "url4": true},
				},
			},
			args: args{serviceName: "service3"},
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

			if got := tt.d.hasService(tt.args.serviceName); got != tt.want {
				t.Errorf("dynamicConfig.hasService() = %v, want %v", got, tt.want)
			}
		})
	}
}
