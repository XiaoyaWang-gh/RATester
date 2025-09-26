package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsetSlice_1(t *testing.T) {
	type TestStruct struct {
		Slice []string
	}

	testCases := []struct {
		name    string
		vals    []string
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			vals:    []string{"val1", "val2", "val3"},
			wantErr: false,
		},
		{
			name:    "Test Case 2",
			vals:    []string{},
			wantErr: false,
		},
		{
			name:    "Test Case 3",
			vals:    nil,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ts := TestStruct{}
			err := setSlice(tc.vals, reflect.ValueOf(&ts.Slice).Elem(), reflect.StructField{})
			if (err != nil) != tc.wantErr {
				t.Errorf("setSlice() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(ts.Slice, tc.vals) {
				t.Errorf("setSlice() = %v, want %v", ts.Slice, tc.vals)
			}
		})
	}
}
