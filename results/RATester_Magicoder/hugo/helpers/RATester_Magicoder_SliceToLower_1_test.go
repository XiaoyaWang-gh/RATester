package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceToLower_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Test case 1",
			args: []string{"Hello", "WORLD"},
			want: []string{"hello", "world"},
		},
		{
			name: "Test case 2",
			args: []string{"Golang", "is", "AWESOME"},
			want: []string{"golang", "is", "awesome"},
		},
		{
			name: "Test case 3",
			args: []string{"123", "456", "789"},
			want: []string{"123", "456", "789"},
		},
		{
			name: "Test case 4",
			args: []string{"", ""},
			want: []string{"", ""},
		},
		{
			name: "Test case 5",
			args: []string{"", "HELLO", ""},
			want: []string{"", "hello", ""},
		},
		{
			name: "Test case 6",
			args: nil,
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

			if got := SliceToLower(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
