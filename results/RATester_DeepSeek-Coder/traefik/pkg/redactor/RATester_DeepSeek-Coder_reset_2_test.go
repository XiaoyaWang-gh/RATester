package redactor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReset_2(t *testing.T) {
	type testStruct struct {
		field1 string
		field2 int
	}

	testCases := []struct {
		name    string
		input   reflect.Value
		wantErr bool
	}{
		{
			name:    "Test case 1: reset a string field",
			input:   reflect.ValueOf("test"),
			wantErr: false,
		},
		{
			name:    "Test case 2: reset a int field",
			input:   reflect.ValueOf(1),
			wantErr: false,
		},
		{
			name:    "Test case 3: reset a struct field",
			input:   reflect.ValueOf(testStruct{"test", 1}),
			wantErr: false,
		},
		{
			name:    "Test case 4: reset a pointer field",
			input:   reflect.ValueOf(&testStruct{"test", 1}),
			wantErr: false,
		},
		{
			name:    "Test case 5: reset a map field",
			input:   reflect.ValueOf(map[string]string{"test": "test"}),
			wantErr: false,
		},
		{
			name:    "Test case 6: reset a slice field",
			input:   reflect.ValueOf([]string{"test"}),
			wantErr: false,
		},
		{
			name:    "Test case 7: reset a interface field",
			input:   reflect.ValueOf(1),
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

			err := reset(tc.input, "test")
			if (err != nil) != tc.wantErr {
				t.Errorf("reset() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
