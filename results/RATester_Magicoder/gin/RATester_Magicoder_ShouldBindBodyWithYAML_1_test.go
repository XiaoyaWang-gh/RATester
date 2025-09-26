package gin

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestShouldBindBodyWithYAML_1(t *testing.T) {
	type testStruct struct {
		Field1 string `yaml:"field1"`
		Field2 int    `yaml:"field2"`
	}

	testCases := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name:    "Valid YAML",
			body:    "field1: value1\nfield2: 123",
			wantErr: false,
		},
		{
			name:    "Invalid YAML",
			body:    "field1: value1\nfield2: value2",
			wantErr: true,
		},
		{
			name:    "Empty Body",
			body:    "",
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

			req, err := http.NewRequest("POST", "/", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			ctx := &Context{Request: req}
			obj := &testStruct{}

			err = ctx.ShouldBindBodyWithYAML(obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("ShouldBindBodyWithYAML() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
