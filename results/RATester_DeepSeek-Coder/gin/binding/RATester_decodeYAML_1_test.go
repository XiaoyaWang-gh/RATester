package binding

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDecodeYAML_1(t *testing.T) {
	type testStruct struct {
		Field1 string `yaml:"field1"`
		Field2 int    `yaml:"field2"`
	}

	tests := []struct {
		name    string
		input   string
		want    testStruct
		wantErr bool
	}{
		{
			name: "valid input",
			input: `
field1: test
field2: 123
`,
			want: testStruct{
				Field1: "test",
				Field2: 123,
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			input: `
field1: test
field3: 123
`,
			want: testStruct{
				Field1: "test",
				Field2: 123,
			},
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
			var got testStruct
			err := decodeYAML(r, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeYAML() got = %v, want %v", got, tt.want)
			}
		})
	}
}
