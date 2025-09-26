package converter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewProvider_2(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		create func(ctx DocumentContext) (Converter, error)
		want   Provider
	}{
		{
			name: "Test case 1",
			create: func(ctx DocumentContext) (Converter, error) {
				return nil, nil
			},
			want: newConverter{
				name: "Test case 1",
				create: func(ctx DocumentContext) (Converter, error) {
					return nil, nil
				},
			},
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			got := NewProvider(tc.name, tc.create)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("NewProvider() = %v, want %v", got, tc.want)
			}
		})
	}
}
