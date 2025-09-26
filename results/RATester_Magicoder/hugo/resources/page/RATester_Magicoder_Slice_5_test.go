package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice_5(t *testing.T) {
	type testCase struct {
		name    string
		input   any
		want    any
		wantErr bool
	}

	tests := []testCase{
		{
			name:  "Test Case 1",
			input: WeightedPages{},
			want:  WeightedPages{},
		},
		{
			name:  "Test Case 2",
			input: []any{},
			want:  WeightedPages{},
		},
		{
			name:  "Test Case 3",
			input: []any{WeightedPage{}, WeightedPage{}},
			want:  WeightedPages{WeightedPage{}, WeightedPage{}},
		},
		{
			name:    "Test Case 4",
			input:   []any{"invalid"},
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

			got, err := WeightedPage{}.Slice(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
