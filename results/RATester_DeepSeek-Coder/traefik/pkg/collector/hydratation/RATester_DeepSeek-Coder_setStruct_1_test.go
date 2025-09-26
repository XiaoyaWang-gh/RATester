package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetStruct_1(t *testing.T) {
	type testStruct struct {
		field1 string
		field2 int
	}

	testCases := []struct {
		name    string
		input   testStruct
		wantErr bool
	}{
		{
			name: "success",
			input: testStruct{
				field1: "test",
				field2: 1,
			},
			wantErr: false,
		},
		{
			name: "failure",
			input: testStruct{
				field1: "test",
				field2: 1,
			},
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

			v := reflect.ValueOf(&tc.input).Elem()
			err := setStruct(v)
			if (err != nil) != tc.wantErr {
				t.Errorf("setStruct() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
