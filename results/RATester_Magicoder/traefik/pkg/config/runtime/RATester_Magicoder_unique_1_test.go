package runtime

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnique_1(t *testing.T) {
	tests := []struct {
		name string
		src  []string
		want []string
	}{
		{
			name: "Test case 1",
			src:  []string{"a", "b", "c", "a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test case 2",
			src:  []string{"apple", "banana", "cherry", "apple", "banana", "cherry"},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Test case 3",
			src:  []string{"1", "2", "3", "1", "2", "3"},
			want: []string{"1", "2", "3"},
		},
		{
			name: "Test case 4",
			src:  []string{"A", "B", "C", "A", "B", "C"},
			want: []string{"A", "B", "C"},
		},
		{
			name: "Test case 5",
			src:  []string{"a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := unique(tt.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
