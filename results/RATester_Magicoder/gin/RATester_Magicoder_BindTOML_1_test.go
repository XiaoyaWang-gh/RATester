package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestBindTOML_1(t *testing.T) {
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
			name: "Success",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`field1 = "value1"
field2 = 123`)),
				},
			},
			obj: &testStruct{},
		},
		{
			name: "Error",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`field1 = "value1"
field2 = "not_an_int"`)),
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

			err := tc.context.BindTOML(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindTOML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
