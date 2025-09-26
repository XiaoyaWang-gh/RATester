package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueStringsReuse_1(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{
			name: "Test case 1",
			s:    []string{"a", "b", "a", "c", "b", "d"},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "Test case 2",
			s:    []string{"apple", "orange", "apple", "banana", "orange", "grape"},
			want: []string{"apple", "orange", "banana", "grape"},
		},
		{
			name: "Test case 3",
			s:    []string{"1", "2", "1", "3", "2", "4"},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "Test case 4",
			s:    []string{"A", "B", "A", "C", "B", "D"},
			want: []string{"A", "B", "C", "D"},
		},
		{
			name: "Test case 5",
			s:    []string{"", "", "a", "b", "", "a"},
			want: []string{"", "a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := UniqueStringsReuse(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStringsReuse() = %v, want %v", got, tt.want)
			}
		})
	}
}
