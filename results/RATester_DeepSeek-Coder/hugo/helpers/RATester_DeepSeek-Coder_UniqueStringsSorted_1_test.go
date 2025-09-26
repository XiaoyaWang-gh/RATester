package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueStringsSorted_1(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{
			name: "Test case 1",
			s:    []string{"apple", "banana", "apple", "pear", "banana", "orange"},
			want: []string{"apple", "banana", "orange", "pear"},
		},
		{
			name: "Test case 2",
			s:    []string{"a", "b", "a", "b", "c", "d", "c"},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "Test case 3",
			s:    []string{"1", "2", "1", "3", "2", "4", "3"},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "Test case 4",
			s:    []string{},
			want: nil,
		},
		{
			name: "Test case 5",
			s:    []string{"onlyone"},
			want: []string{"onlyone"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := UniqueStringsSorted(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStringsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
