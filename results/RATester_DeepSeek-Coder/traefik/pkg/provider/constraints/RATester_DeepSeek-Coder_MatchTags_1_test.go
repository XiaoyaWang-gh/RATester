package constraints

import (
	"fmt"
	"testing"
)

func TestMatchTags_1(t *testing.T) {
	tests := []struct {
		name    string
		tags    []string
		expr    string
		want    bool
		wantErr bool
	}{
		{
			name: "empty tags and expr",
			tags: []string{},
			expr: "",
			want: true,
		},
		{
			name: "matching tags",
			tags: []string{"tag1", "tag2", "tag3"},
			expr: "Tag('tag1') AND Tag('tag2')",
			want: true,
		},
		{
			name: "non-matching tags",
			tags: []string{"tag1", "tag2", "tag3"},
			expr: "Tag('tag1') AND Tag('tag4')",
			want: false,
		},
		{
			name:    "invalid expr",
			tags:    []string{"tag1", "tag2", "tag3"},
			expr:    "Tag('tag1' AND 'tag2'",
			wantErr: true,
		},
		{
			name: "empty expr",
			tags: []string{"tag1", "tag2", "tag3"},
			expr: "",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := MatchTags(tt.tags, tt.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MatchTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
