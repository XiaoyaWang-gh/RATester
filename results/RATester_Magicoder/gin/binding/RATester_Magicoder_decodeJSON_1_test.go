package binding

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestdecodeJSON_1(t *testing.T) {
	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name    string
		input   string
		want    testStruct
		wantErr bool
	}{
		{
			name:  "valid JSON",
			input: `{"field1": "value1", "field2": 123}`,
			want:  testStruct{Field1: "value1", Field2: 123},
		},
		{
			name:    "invalid JSON",
			input:   `{"field1": "value1", "field2": "not a number"}`,
			wantErr: true,
		},
		{
			name:    "missing field",
			input:   `{"field1": "value1"}`,
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

			var got testStruct
			err := decodeJSON(strings.NewReader(tc.input), &got)
			if (err != nil) != tc.wantErr {
				t.Errorf("decodeJSON() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("decodeJSON() = %v, want %v", got, tc.want)
			}
		})
	}
}
