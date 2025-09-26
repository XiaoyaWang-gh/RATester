package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecode_1(t *testing.T) {
	type TestStruct struct {
		Field1 string
		Field2 int
	}

	testCases := []struct {
		name    string
		input   map[string][]string
		output  TestStruct
		wantErr bool
	}{
		{
			name: "success",
			input: map[string][]string{
				"Field1": {"value1"},
				"Field2": {"100"},
			},
			output: TestStruct{
				Field1: "value1",
				Field2: 100,
			},
			wantErr: false,
		},
		{
			name: "missing value",
			input: map[string][]string{
				"Field1": {"value1"},
			},
			output:  TestStruct{},
			wantErr: true,
		},
		{
			name: "invalid value",
			input: map[string][]string{
				"Field1": {"value1"},
				"Field2": {"not_an_int"},
			},
			output:  TestStruct{},
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

			d := &Decoder{}
			var v TestStruct
			err := d.decode(reflect.ValueOf(&v).Elem(), "", []pathPart{{}}, tc.input["Field1"])
			if (err != nil) != tc.wantErr {
				t.Errorf("decode() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(v, tc.output) {
				t.Errorf("decode() = %v, want %v", v, tc.output)
			}
		})
	}
}
