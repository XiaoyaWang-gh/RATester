package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func Testencode_1(t *testing.T) {
	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name    string
		input   any
		want    map[string][]string
		wantErr bool
	}{
		{
			name: "valid struct",
			input: testStruct{
				Field1: "value1",
				Field2: 123,
			},
			want: map[string][]string{
				"field1": {"value1"},
				"field2": {"123"},
			},
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   "invalid",
			want:    nil,
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

			e := &Encoder{}
			dst := make(map[string][]string)
			err := e.encode(reflect.ValueOf(tc.input), dst)

			if (err != nil) != tc.wantErr {
				t.Errorf("encode() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(dst, tc.want) {
				t.Errorf("encode() = %v, want %v", dst, tc.want)
			}
		})
	}
}
