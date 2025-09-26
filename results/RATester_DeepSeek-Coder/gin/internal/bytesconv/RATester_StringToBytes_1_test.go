package bytesconv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToBytes_1(t *testing.T) {
	testCases := []struct {
		name string
		arg  string
		want []byte
	}{
		{
			name: "Test case 1",
			arg:  "Hello, world",
			want: []byte("Hello, world"),
		},
		{
			name: "Test case 2",
			arg:  "Golang",
			want: []byte("Golang"),
		},
		{
			name: "Test case 3",
			arg:  "1234567890",
			want: []byte("1234567890"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := StringToBytes(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("StringToBytes(%s) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}
