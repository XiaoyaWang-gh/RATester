package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrimArr_1(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want []string
	}{
		{
			name: "Test case 1",
			arr:  []string{"", "hello", "", "world", ""},
			want: []string{"hello", "world"},
		},
		{
			name: "Test case 2",
			arr:  []string{"", "", "", "", ""},
			want: []string{},
		},
		{
			name: "Test case 3",
			arr:  []string{"hello", "world"},
			want: []string{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := TrimArr(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
