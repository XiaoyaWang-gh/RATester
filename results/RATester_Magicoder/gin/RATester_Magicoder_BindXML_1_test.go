package gin

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBindXML_1(t *testing.T) {
	type testStruct struct {
		Field1 string `xml:"field1"`
		Field2 int    `xml:"field2"`
	}

	testCases := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name:    "valid XML",
			body:    `<testStruct><field1>value1</field1><field2>123</field2></testStruct>`,
			wantErr: false,
		},
		{
			name:    "invalid XML",
			body:    `<testStruct><field1>value1</field1><field2>abc</field2></testStruct>`,
			wantErr: true,
		},
		{
			name:    "empty body",
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
			req.Header.Set("Content-Type", "application/xml")

			ctx := &Context{Request: req}
			obj := &testStruct{}

			err = ctx.BindXML(obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindXML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
