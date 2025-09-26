package hashing

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cespare/xxhash/v2"
)

func TestGetXxHashReadFrom_1(t *testing.T) {
	testCases := []struct {
		name string
		want *xxhashReadFrom
	}{
		{
			name: "Test case 1",
			want: &xxhashReadFrom{
				buff:   make([]byte, 1024),
				Digest: xxhash.New(),
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := getXxHashReadFrom()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("getXxHashReadFrom() = %v, want %v", got, tc.want)
			}
		})
	}
}
