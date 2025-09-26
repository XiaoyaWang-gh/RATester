package data

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttoURLAndHeaders_1(t *testing.T) {
	tests := []struct {
		name        string
		urlParts    []any
		wantURL     string
		wantHeaders map[string]any
	}{
		{
			name:        "Test case 1",
			urlParts:    []any{"http://example.com", map[string]any{"field name": "value"}},
			wantURL:     "http://example.com",
			wantHeaders: map[string]any{"field name": "value"},
		},
		{
			name:        "Test case 2",
			urlParts:    []any{"http://example.com", "value"},
			wantURL:     "http://example.com",
			wantHeaders: nil,
		},
		{
			name:        "Test case 3",
			urlParts:    []any{"http://example.com"},
			wantURL:     "",
			wantHeaders: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotURL, gotHeaders := toURLAndHeaders(tt.urlParts)
			if gotURL != tt.wantURL {
				t.Errorf("toURLAndHeaders() gotURL = %v, want %v", gotURL, tt.wantURL)
			}
			if !reflect.DeepEqual(gotHeaders, tt.wantHeaders) {
				t.Errorf("toURLAndHeaders() gotHeaders = %v, want %v", gotHeaders, tt.wantHeaders)
			}
		})
	}
}
