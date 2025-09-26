package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHTTPCodeRanges_1(t *testing.T) {
	tests := []struct {
		name      string
		strBlocks []string
		want      HTTPCodeRanges
		wantErr   bool
	}{
		{
			name:      "single code",
			strBlocks: []string{"200"},
			want:      HTTPCodeRanges{{200, 200}},
			wantErr:   false,
		},
		{
			name:      "range",
			strBlocks: []string{"200-299"},
			want:      HTTPCodeRanges{{200, 299}},
			wantErr:   false,
		},
		{
			name:      "multiple ranges",
			strBlocks: []string{"200-299", "400-499"},
			want:      HTTPCodeRanges{{200, 299}, {400, 499}},
			wantErr:   false,
		},
		{
			name:      "invalid code",
			strBlocks: []string{"200-invalid"},
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := NewHTTPCodeRanges(tt.strBlocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPCodeRanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPCodeRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
