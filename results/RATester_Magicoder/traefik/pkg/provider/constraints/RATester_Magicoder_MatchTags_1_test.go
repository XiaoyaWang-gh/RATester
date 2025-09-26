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
			name:    "empty expr",
			tags:    []string{"tag1", "tag2"},
			expr:    "",
			want:    true,
			wantErr: false,
		},
		{
			name:    "valid expr",
			tags:    []string{"tag1", "tag2"},
			expr:    "Tag('tag1')",
			want:    true,
			wantErr: false,
		},
		{
			name:    "invalid expr",
			tags:    []string{"tag1", "tag2"},
			expr:    "Tag('tag3')",
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid expr",
			tags:    []string{"tag1", "tag2"},
			expr:    "Tag('tag3')",
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid expr",
			tags:    []string{"tag1", "tag2"},
			expr:    "Tag('tag3')",
			want:    false,
			wantErr: false,
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
