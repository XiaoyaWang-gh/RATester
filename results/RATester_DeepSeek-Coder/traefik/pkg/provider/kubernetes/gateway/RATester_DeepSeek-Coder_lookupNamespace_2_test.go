package gateway

import (
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLookupNamespace_2(t *testing.T) {
	tests := []struct {
		name           string
		isNamespaceAll bool
		namespace      string
		want           string
	}{
		{
			name:           "Test lookupNamespace with isNamespaceAll true",
			isNamespaceAll: true,
			namespace:      "test",
			want:           metav1.NamespaceAll,
		},
		{
			name:           "Test lookupNamespace with isNamespaceAll false",
			isNamespaceAll: false,
			namespace:      "test",
			want:           "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &clientWrapper{
				isNamespaceAll: tt.isNamespaceAll,
			}
			got := c.lookupNamespace(tt.namespace)
			if got != tt.want {
				t.Errorf("lookupNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
