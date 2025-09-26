package binding

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestdecodeToml_1(t *testing.T) {
	type testStruct struct {
		Field1 string
		Field2 int
	}

	testCases := []struct {
		name    string
		input   string
		want    testStruct
		wantErr bool
	}{
		{
			name:  "valid toml",
			input: `Field1 = "value1" Field2 = 123`,
			want:  testStruct{Field1: "value1", Field2: 123},
		},
		{
			name:    "invalid toml",
			input:   `Field1 = "value1" Field2 = "abc"`,
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

			r := strings.NewReader(tc.input)
			var got testStruct
			err := decodeToml(r, &got)
			if (err != nil) != tc.wantErr {
				t.Errorf("decodeToml() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("decodeToml() = %v, want %v", got, tc.want)
			}
		})
	}
}
