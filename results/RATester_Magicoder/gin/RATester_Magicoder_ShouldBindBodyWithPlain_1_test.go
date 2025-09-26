package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestShouldBindBodyWithPlain_1(t *testing.T) {
	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
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
					Body: ioutil.NopCloser(strings.NewReader(`{"field1": "value1", "field2": 123}`)),
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "Invalid JSON",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`{"field1": "value1", "field2": "not a number"}`)),
				},
			},
			obj:     &testStruct{},
			wantErr: true,
		},
		{
			name: "Nil Object",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`{"field1": "value1", "field2": 123}`)),
				},
			},
			obj:     nil,
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

			err := tc.context.ShouldBindBodyWithPlain(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("ShouldBindBodyWithPlain() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
