package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetSlice_1(t *testing.T) {
	testCases := []struct {
		name    string
		input   reflect.Value
		wantErr bool
	}{
		{
			name:    "Test case 1: Success",
			input:   reflect.ValueOf([]int{1, 2, 3}),
			wantErr: false,
		},
		{
			name:    "Test case 2: Failure",
			input:   reflect.ValueOf([]int{1, 2, 3}),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := setSlice(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("setSlice() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
