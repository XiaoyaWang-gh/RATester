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
			name:         "Test Case 1",
			mediaType:    media.Type{Type: "application/rss+xml"},
			resourceType: "rss",
			want:         resourceTypesHolder{mediaType: media.Type{Type: "application/rss+xml"}, resourceType: "rss"},
		},
		{
			name:         "Test Case 2",
			mediaType:    media.Type{Type: "application/json"},
			resourceType: "json",
			want:         resourceTypesHolder{mediaType: media.Type{Type: "application/json"}, resourceType: "json"},
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

			if got := NewResourceTypesProvider(tt.mediaType, tt.resourceType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceTypesProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
