package tcp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBraceIndices_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{
			name:    "empty string",
			input:   "",
			want:    nil,
			wantErr: false,
		},
		{
			name:    "single brace",
			input:   "{}",
			want:    []int{0, 2},
			wantErr: false,
		},
		{
			name:    "nested braces",
			input:   "{}{}",
			want:    []int{0, 2, 3, 5},
			wantErr: false,
		},
		{
			name:    "unbalanced braces",
			input:   "{{}",
			want:    nil,
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

			got, err := braceIndices(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("braceIndices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("braceIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
