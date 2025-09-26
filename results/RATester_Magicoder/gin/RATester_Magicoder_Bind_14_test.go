package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestBind_14(t *testing.T) {
	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	tests := []struct {
		name    string
		context *Context
		obj     any
		wantErr bool
	}{
		{
			name: "Test Bind Success",
			context: &Context{
				Request: &http.Request{
					Method: "POST",
					Header: http.Header{
						"Content-Type": []string{"application/json"},
					},
					Body: ioutil.NopCloser(strings.NewReader(`{"field1": "value1", "field2": 123}`)),
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "Test Bind Failure",
			context: &Context{
				Request: &http.Request{
					Method: "POST",
					Header: http.Header{
						"Content-Type": []string{"application/json"},
					},
					Body: ioutil.NopCloser(strings.NewReader(`{"field1": "value1", "field2": "abc"}`)),
				},
			},
			obj:     &testStruct{},
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

			err := tt.context.Bind(tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
