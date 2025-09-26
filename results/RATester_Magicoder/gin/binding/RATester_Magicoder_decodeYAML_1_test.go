package binding

import (
	"fmt"
	"strings"
	"testing"
)

func TestdecodeYAML_1(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid yaml",
			input: `
name: John Doe
age: 30
`,
			wantErr: false,
		},
		{
			name: "invalid yaml",
			input: `
name: John Doe
age: 30
invalid: field
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := strings.NewReader(tt.input)
			var obj struct {
				Name string `yaml:"name"`
				Age  int    `yaml:"age"`
			}
			err := decodeYAML(r, &obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
