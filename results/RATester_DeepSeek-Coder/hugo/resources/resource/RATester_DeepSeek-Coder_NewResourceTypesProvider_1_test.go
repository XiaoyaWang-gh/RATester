package resource

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestNewResourceTypesProvider_1(t *testing.T) {
	tests := []struct {
		name         string
		mediaType    media.Type
		resourceType string
		want         ResourceTypesProvider
	}{
		{
			name:         "Test case 1",
			mediaType:    media.Type{Type: "application/json"},
			resourceType: "test",
			want:         resourceTypesHolder{mediaType: media.Type{Type: "application/json"}, resourceType: "test"},
		},
		{
			name:         "Test case 2",
			mediaType:    media.Type{Type: "application/xml"},
			resourceType: "test2",
			want:         resourceTypesHolder{mediaType: media.Type{Type: "application/xml"}, resourceType: "test2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewResourceTypesProvider(tt.mediaType, tt.resourceType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceTypesProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
