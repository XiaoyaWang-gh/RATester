package pageparser

import (
	"fmt"
	"testing"
)

func TestLexFrontMatterJSON_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name: "valid JSON",
			input: `{"title": "Test Page", "date": "2022-01-01", "tags": ["tag1", "tag2"]}
Some content here...`,
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			input:   `{"title": "Test Page", "date": "2022-01-01", "tags": ["tag1", "tag2"`,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   ``,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &pageLexer{
				input: []byte(tt.input),
			}

			got := lexFrontMatterJSON(l)

			if (got != nil) != tt.wantErr {
				t.Errorf("lexFrontMatterJSON() error = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}
