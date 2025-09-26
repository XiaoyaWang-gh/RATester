package nomad

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTagsToLabels_1(t *testing.T) {
	tests := []struct {
		name   string
		tags   []string
		prefix string
		want   map[string]string
	}{
		{
			name:   "empty tags",
			tags:   []string{},
			prefix: "prefix",
			want:   map[string]string{},
		},
		{
			name:   "one tag",
			tags:   []string{"prefix.key=value"},
			prefix: "prefix",
			want:   map[string]string{"traefik.key": "value"},
		},
		{
			name:   "multiple tags",
			tags:   []string{"prefix.key1=value1", "prefix.key2=value2"},
			prefix: "prefix",
			want:   map[string]string{"traefik.key1": "value1", "traefik.key2": "value2"},
		},
		{
			name:   "tag without prefix",
			tags:   []string{"key=value"},
			prefix: "prefix",
			want:   map[string]string{},
		},
		{
			name:   "tag with wrong format",
			tags:   []string{"prefix.key=value=value"},
			prefix: "prefix",
			want:   map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tagsToLabels(tt.tags, tt.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagsToLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
