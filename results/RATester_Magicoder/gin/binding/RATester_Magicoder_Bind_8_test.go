package binding

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBind_8(t *testing.T) {
	type testStruct struct {
		Field1 string `xml:"field1"`
		Field2 int    `xml:"field2"`
	}

	testCases := []struct {
		name    string
		body    string
		obj     any
		wantErr bool
	}{
		{
			name: "Valid XML",
			body: `<testStruct><field1>value1</field1><field2>123</field2></testStruct>`,
			obj:  &testStruct{},
		},
		{
			name:    "Invalid XML",
			body:    `<testStruct><field1>value1</field1><field2>abc</field2></testStruct>`,
			obj:     &testStruct{},
			wantErr: true,
		},
		{
			name: "Empty XML",
			body: ``,
			obj:  &testStruct{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))
			err := xmlBinding{}.Bind(req, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
