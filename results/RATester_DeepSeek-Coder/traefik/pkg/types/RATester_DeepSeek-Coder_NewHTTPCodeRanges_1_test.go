package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHTTPCodeRanges_1(t *testing.T) {
	tests := []struct {
		name       string
		strBlocks  []string
		wantBlocks HTTPCodeRanges
		wantErr    bool
	}{
		{
			name:      "single code",
			strBlocks: []string{"200"},
			wantBlocks: HTTPCodeRanges{
				{200, 200},
			},
			wantErr: false,
		},
		{
			name:      "range",
			strBlocks: []string{"200-299"},
			wantBlocks: HTTPCodeRanges{
				{200, 299},
			},
			wantErr: false,
		},
		{
			name:      "multiple ranges",
			strBlocks: []string{"200-299", "400-450"},
			wantBlocks: HTTPCodeRanges{
				{200, 299},
				{400, 450},
			},
			wantErr: false,
		},
		{
			name:       "invalid code",
			strBlocks:  []string{"200-invalid"},
			wantBlocks: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotBlocks, err := NewHTTPCodeRanges(tt.strBlocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPCodeRanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBlocks, tt.wantBlocks) {
				t.Errorf("NewHTTPCodeRanges() = %v, want %v", gotBlocks, tt.wantBlocks)
			}
		})
	}
}
