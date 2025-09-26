package layouts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestuniqueStringsReuse_2(t *testing.T) {
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
			s:    []string{"apple", "banana", "apple", "cherry", "banana", "date"},
			want: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name: "Test case 3",
			s:    []string{"1", "2", "1", "3", "2", "4"},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "Test case 4",
			s:    []string{"hello", "world", "hello", "golang", "world", "test"},
			want: []string{"hello", "world", "golang", "test"},
		},
		{
			name: "Test case 5",
			s:    []string{"a", "a", "a", "a", "a", "a"},
			want: []string{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := uniqueStringsReuse(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueStringsReuse() = %v, want %v", got, tt.want)
			}
		})
	}
}
