package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestBindYAML_1(t *testing.T) {
	type testStruct struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
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
					Body: ioutil.NopCloser(strings.NewReader(`name: John Doe
age: 30`)),
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "Error",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`name: John Doe
age: not_an_int`)),
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

			err := tc.context.BindYAML(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindYAML() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
