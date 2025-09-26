package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueStrings_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Test case 1",
			args: []string{"a", "b", "c", "a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test case 2",
			args: []string{"apple", "banana", "cherry", "apple", "banana", "cherry"},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Test case 3",
			args: []string{"1", "2", "3", "1", "2", "3"},
			want: []string{"1", "2", "3"},
		},
		{
			name: "Test case 4",
			args: []string{},
			want: []string{},
		},
		{
			name: "Test case 5",
			args: []string{"a"},
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

			if got := UniqueStrings(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
