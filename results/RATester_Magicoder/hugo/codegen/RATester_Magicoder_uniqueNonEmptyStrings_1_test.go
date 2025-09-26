package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestuniqueNonEmptyStrings_1(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{
			name: "Test case 1",
			s:    []string{"a", "b", "a", "", "c", "", "d"},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "Test case 2",
			s:    []string{"", "", "", "", "", ""},
			want: []string{},
		},
		{
			name: "Test case 3",
			s:    []string{"a", "a", "a", "a", "a", "a"},
			want: []string{"a"},
		},
		{
			name: "Test case 4",
			s:    []string{"", "b", "", "b", "", "b", ""},
			want: []string{"b"},
		},
		{
			name: "Test case 5",
			s:    []string{"a", "b", "c", "d", "e", "f"},
			want: []string{"a", "b", "c", "d", "e", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := uniqueNonEmptyStrings(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueNonEmptyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
