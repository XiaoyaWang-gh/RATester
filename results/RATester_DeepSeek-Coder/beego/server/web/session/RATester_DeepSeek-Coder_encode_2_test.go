package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncode_2(t *testing.T) {
	testCases := []struct {
		name  string
		value []byte
		want  []byte
	}{
		{
			name:  "Test case 1",
			value: []byte("test"),
			want:  []byte("dGVzdA=="),
		},
		{
			name:  "Test case 2",
			value: []byte("hello"),
			want:  []byte("aGVsbG8="),
		},
		{
			name:  "Test case 3",
			value: []byte("world"),
			want:  []byte("d29ybGQ="),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := encode(tc.value)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("encode(%v) = %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}
