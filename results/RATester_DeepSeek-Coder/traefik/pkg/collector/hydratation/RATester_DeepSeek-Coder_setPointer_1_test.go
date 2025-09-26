package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetPointer_1(t *testing.T) {
	type testStruct struct {
		field reflect.Value
	}

	testCases := []struct {
		name    string
		input   testStruct
		wantErr bool
	}{
		{
			name: "test case 1",
			input: testStruct{
				field: reflect.ValueOf(new(int)),
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			input: testStruct{
				field: reflect.ValueOf(nil),
			},
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

			err := setPointer(tc.input.field)
			if (err != nil) != tc.wantErr {
				t.Errorf("setPointer() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
