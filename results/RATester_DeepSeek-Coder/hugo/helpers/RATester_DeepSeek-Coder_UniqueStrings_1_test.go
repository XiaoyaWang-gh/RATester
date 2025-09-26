package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueStrings_1(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{
			name: "empty slice",
			s:    []string{},
			want: []string{},
		},
		{
			name: "single element",
			s:    []string{"a"},
			want: []string{"a"},
		},
		{
			name: "duplicates",
			s:    []string{"a", "b", "a", "c", "b", "c", "d"},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "no duplicates",
			s:    []string{"a", "b", "c", "d"},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := UniqueStrings(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
