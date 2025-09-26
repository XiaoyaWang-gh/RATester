package tcp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBraceIndices_1(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    []int
		wantErr bool
	}{
		{
			name:    "empty string",
			s:       "",
			want:    nil,
			wantErr: false,
		},
		{
			name:    "unbalanced braces",
			s:       "{",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "balanced braces",
			s:       "{test}",
			want:    []int{0, 6},
			wantErr: false,
		},
		{
			name:    "nested braces",
			s:       "{test{nested}}",
			want:    []int{0, 16},
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

			got, err := braceIndices(tt.s)
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
