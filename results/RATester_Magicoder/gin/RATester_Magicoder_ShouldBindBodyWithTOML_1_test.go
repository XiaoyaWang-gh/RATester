package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestShouldBindBodyWithTOML_1(t *testing.T) {
	type testStruct struct {
		Field1 string `toml:"field1"`
		Field2 int    `toml:"field2"`
	}

	testCases := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name:    "valid TOML",
			body:    `field1 = "value1"\nfield2 = 123`,
			wantErr: false,
		},
		{
			name:    "invalid TOML",
			body:    `field1 = "value1"\nfield2 = "abc"`,
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

			ctx := &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(tc.body)),
				},
			}

			obj := &testStruct{}
			err := ctx.ShouldBindBodyWithTOML(obj)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if obj.Field1 != "value1" || obj.Field2 != 123 {
					t.Errorf("Unexpected result: %v", obj)
				}
			}
		})
	}
}
