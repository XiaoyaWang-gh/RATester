package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	kschema "k8s.io/apimachinery/pkg/runtime/schema"
)

func TestResource_1(t *testing.T) {
	tests := []struct {
		name     string
		resource string
		want     kschema.GroupResource
	}{
		{
			name:     "Test case 1",
			resource: "test",
			want:     kschema.GroupResource{Group: "", Resource: "test"},
		},
		{
			name:     "Test case 2",
			resource: "test2",
			want:     kschema.GroupResource{Group: "", Resource: "test2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Resource(tt.resource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resource() = %v, want %v", got, tt.want)
			}
		})
	}
}
