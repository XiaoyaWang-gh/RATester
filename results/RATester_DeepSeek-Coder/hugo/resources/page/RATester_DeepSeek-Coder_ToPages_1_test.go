package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToPages_1(t *testing.T) {
	type testCase struct {
		name    string
		input   any
		want    Pages
		wantErr bool
	}

	tests := []testCase{
		{
			name:  "nil",
			input: nil,
			want:  Pages{},
		},
		{
			name:  "Pages",
			input: Pages{},
			want:  Pages{},
		},
		{
			name:  "*Pages",
			input: &Pages{},
			want:  Pages{},
		},
		{
			name:  "WeightedPages",
			input: WeightedPages{},
			want:  Pages{},
		},
		{
			name:  "PageGroup",
			input: PageGroup{},
			want:  Pages{},
		},
		{
			name:  "[]Page",
			input: []Page{},
			want:  Pages{},
		},
		{
			name:  "[]any",
			input: []any{},
			want:  Pages{},
		},
		{
			name:    "unsupported type",
			input:   123,
			want:    Pages{},
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

			got, err := ToPages(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPages() = %v, want %v", got, tt.want)
			}
		})
	}
}
