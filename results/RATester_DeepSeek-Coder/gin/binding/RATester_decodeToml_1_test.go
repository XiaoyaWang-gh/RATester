package binding

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDecodeToml_1(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name    string
		input   string
		want    testStruct
		wantErr bool
	}{
		{
			name:  "valid toml",
			input: "Name = 'John Doe'\nAge = 30",
			want:  testStruct{Name: "John Doe", Age: 30},
		},
		{
			name:    "invalid toml",
			input:   "Name = 'John Doe'\nAge = 'thirty'",
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
				t.Errorf("decodeToml() got = %v, want %v", got, tc.want)
			}
		})
	}
}
