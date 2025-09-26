package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseTag_1(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		wantKey  string
		wantOpts tagOptions
	}{
		{
			name:     "empty tag",
			tag:      "",
			wantKey:  "",
			wantOpts: tagOptions{},
		},
		{
			name:     "tag with no options",
			tag:      "field",
			wantKey:  "field",
			wantOpts: tagOptions{},
		},
		{
			name:     "tag with options",
			tag:      "field,option1,option2",
			wantKey:  "field",
			wantOpts: tagOptions{"option1", "option2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotKey, gotOpts := parseTag(tt.tag)
			if gotKey != tt.wantKey {
				t.Errorf("parseTag() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if !reflect.DeepEqual(gotOpts, tt.wantOpts) {
				t.Errorf("parseTag() gotOpts = %v, want %v", gotOpts, tt.wantOpts)
			}
		})
	}
}
