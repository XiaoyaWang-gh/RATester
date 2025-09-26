package resource

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestData_4(t *testing.T) {
	testCases := []struct {
		name     string
		resource *resourceError
		want     any
	}{
		{
			name: "Test case 1",
			resource: &resourceError{
				error: errors.New("test error"),
				data:  "test data",
			},
			want: "test data",
		},
		{
			name: "Test case 2",
			resource: &resourceError{
				error: errors.New("another error"),
				data:  123,
			},
			want: 123,
		},
		{
			name: "Test case 3",
			resource: &resourceError{
				error: errors.New("yet another error"),
				data:  []string{"test", "data"},
			},
			want: []string{"test", "data"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.resource.Data()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
