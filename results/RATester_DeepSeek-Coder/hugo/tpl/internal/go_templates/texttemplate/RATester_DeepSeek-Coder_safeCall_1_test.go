package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSafeCall_1(t *testing.T) {
	testCases := []struct {
		name    string
		fun     reflect.Value
		args    []reflect.Value
		wantVal reflect.Value
		wantErr error
	}{
		{
			name:    "test safeCall with no error",
			fun:     reflect.ValueOf(func(a int, b string) (int, error) { return a + len(b), nil }),
			args:    []reflect.Value{reflect.ValueOf(5), reflect.ValueOf("hello")},
			wantVal: reflect.ValueOf(10),
			wantErr: nil,
		},
		{
			name:    "test safeCall with error",
			fun:     reflect.ValueOf(func(a int, b string) (int, error) { return 0, fmt.Errorf("error") }),
			args:    []reflect.Value{reflect.ValueOf(5), reflect.ValueOf("hello")},
			wantVal: reflect.ValueOf(0),
			wantErr: fmt.Errorf("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotVal, gotErr := safeCall(tc.fun, tc.args)
			if !reflect.DeepEqual(gotVal, tc.wantVal) || !reflect.DeepEqual(gotErr, tc.wantErr) {
				t.Errorf("safeCall() = %v, %v; want %v, %v", gotVal, gotErr, tc.wantVal, tc.wantErr)
			}
		})
	}
}
