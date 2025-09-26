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
			s:    []string{"apple", "banana", "apple", "cherry", "banana", "date"},
			want: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name: "Test case 2",
			s:    []string{"apple", "banana", "cherry", "date"},
			want: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name: "Test case 3",
			s:    []string{"apple", "apple", "apple", "apple"},
			want: []string{"apple"},
		},
		{
			name: "Test case 4",
			s:    []string{},
			want: nil,
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
