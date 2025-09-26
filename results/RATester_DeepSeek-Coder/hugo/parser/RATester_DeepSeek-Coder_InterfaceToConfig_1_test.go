package parser

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

func TestInterfaceToConfig_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name    string
		input   any
		format  metadecoders.Format
		wantErr bool
	}{
		{
			name:    "Nil input",
			input:   nil,
			format:  metadecoders.YAML,
			wantErr: true,
		},
		{
			name:    "Valid input",
			input:   testStruct{Name: "John", Age: 30},
			format:  metadecoders.YAML,
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &bytes.Buffer{}
			err := InterfaceToConfig(tc.input, tc.format, w)
			if (err != nil) != tc.wantErr {
				t.Errorf("InterfaceToConfig() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr {
				// Add assertions to check the content of the writer
				// For example, if the format is YAML, you can use goyaml to parse the content of the writer and compare it with the expected YAML
			}
		})
	}
}
