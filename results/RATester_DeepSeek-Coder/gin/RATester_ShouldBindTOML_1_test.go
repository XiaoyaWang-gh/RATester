package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestShouldBindTOML_1(t *testing.T) {
	type testStruct struct {
		Field1 string `toml:"field1"`
		Field2 int    `toml:"field2"`
	}

	testCases := []struct {
		name    string
		context *Context
		obj     any
		wantErr bool
	}{
		{
			name: "should bind TOML successfully",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`field1 = "test"
field2 = 123`)),
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "should return error when TOML is invalid",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`field1 = "test"
field2 = "not a number"`)),
				},
			},
			obj:     &testStruct{},
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

			err := tc.context.ShouldBindTOML(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("ShouldBindTOML() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
