package template

import (
	"fmt"
	"testing"
)

func TesteatTagName_1(t *testing.T) {
	tests := []struct {
		name  string
		s     []byte
		i     int
		want  int
		want1 element
	}{
		{
			name:  "Test case 1",
			s:     []byte("<div>"),
			i:     1,
			want:  4,
			want1: elementNone,
		},
		{
			name:  "Test case 2",
			s:     []byte("<div->"),
			i:     1,
			want:  4,
			want1: elementNone,
		},
		{
			name:  "Test case 3",
			s:     []byte("<div:>"),
			i:     1,
			want:  4,
			want1: elementNone,
		},
		{
			name:  "Test case 4",
			s:     []byte("<div>"),
			i:     1,
			want:  4,
			want1: elementNone,
		},
		{
			name:  "Test case 5",
			s:     []byte("<div>"),
			i:     1,
			want:  4,
			want1: elementNone,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := eatTagName(tt.s, tt.i)
			if got != tt.want {
				t.Errorf("eatTagName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("eatTagName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
