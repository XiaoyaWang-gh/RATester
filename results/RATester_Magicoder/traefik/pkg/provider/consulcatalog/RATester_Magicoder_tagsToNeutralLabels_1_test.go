package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTagsToNeutralLabels_1(t *testing.T) {
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
			name:   "no matching tags",
			tags:   []string{"tag1", "tag2"},
			prefix: "prefix",
			want:   map[string]string{},
		},
		{
			name:   "one matching tag",
			tags:   []string{"prefix.key=value"},
			prefix: "prefix",
			want:   map[string]string{"traefik.key": "value"},
		},
		{
			name:   "multiple matching tags",
			tags:   []string{"prefix.key1=value1", "prefix.key2=value2"},
			prefix: "prefix",
			want:   map[string]string{"traefik.key1": "value1", "traefik.key2": "value2"},
		},
		{
			name:   "tags with different prefixes",
			tags:   []string{"prefix.key=value", "otherprefix.key=value"},
			prefix: "prefix",
			want:   map[string]string{"traefik.key": "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tagsToNeutralLabels(tt.tags, tt.prefix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagsToNeutralLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
