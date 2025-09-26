package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSafeConvert_1(t *testing.T) {
	type testCase struct {
		name    string
		value   reflect.Value
		t       reflect.Type
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "Test safeConvert with int value and int type",
			value:   reflect.ValueOf(1),
			t:       reflect.TypeOf(1),
			wantErr: false,
		},
		{
			name:    "Test safeConvert with float value and int type",
			value:   reflect.ValueOf(1.1),
			t:       reflect.TypeOf(1),
			wantErr: true,
		},
		{
			name:    "Test safeConvert with string value and string type",
			value:   reflect.ValueOf("test"),
			t:       reflect.TypeOf("test"),
			wantErr: false,
		},
		{
			name:    "Test safeConvert with bool value and bool type",
			value:   reflect.ValueOf(true),
			t:       reflect.TypeOf(true),
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

			_, err := safeConvert(tc.value, tc.t)
			if (err != nil) != tc.wantErr {
				t.Errorf("safeConvert() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
