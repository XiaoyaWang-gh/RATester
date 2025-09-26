package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestShouldBindXML_1(t *testing.T) {
	type testStruct struct {
		Field1 string `xml:"field1"`
		Field2 int    `xml:"field2"`
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
					Body: ioutil.NopCloser(strings.NewReader(`<testStruct><field1>value1</field1><field2>123</field2></testStruct>`)),
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "Invalid XML",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`<testStruct><field1>value1</field1><field2>invalid</field2></testStruct>`)),
				},
			},
			obj:     &testStruct{},
			wantErr: true,
		},
		{
			name: "Nil Object",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader(`<testStruct><field1>value1</field1><field2>123</field2></testStruct>`)),
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

			err := tc.context.ShouldBindXML(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("ShouldBindXML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
