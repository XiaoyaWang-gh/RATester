package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceRand_1(t *testing.T) {
	testCases := []struct {
		name string
		arg  []interface{}
		want interface{}
	}{
		{
			name: "Test case 1",
			arg:  []interface{}{1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "Test case 2",
			arg:  []interface{}{"a", "b", "c", "d", "e"},
			want: "a",
		},
		{
			name: "Test case 3",
			arg:  []interface{}{true, false},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := SliceRand(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SliceRand(%v) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}
