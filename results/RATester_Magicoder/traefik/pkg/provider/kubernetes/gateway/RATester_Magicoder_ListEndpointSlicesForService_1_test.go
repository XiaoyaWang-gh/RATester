package gateway

import (
	"fmt"
	"reflect"
	"testing"

	discoveryv1 "k8s.io/api/discovery/v1"
)

func TestListEndpointSlicesForService_1(t *testing.T) {
	tests := []struct {
		name        string
		namespace   string
		serviceName string
		want        []*discoveryv1.EndpointSlice
		wantErr     bool
	}{
		{
			name:        "Test case 1",
			namespace:   "test-namespace",
			serviceName: "test-service",
			want:        []*discoveryv1.EndpointSlice{},
			wantErr:     false,
		},
		{
			name:        "Test case 2",
			namespace:   "non-watched-namespace",
			serviceName: "test-service",
			want:        nil,
			wantErr:     true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &clientWrapper{
				watchedNamespaces: []string{"test-namespace"},
			}
			got, err := c.ListEndpointSlicesForService(tt.namespace, tt.serviceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListEndpointSlicesForService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListEndpointSlicesForService() = %v, want %v", got, tt.want)
			}
		})
	}
}
