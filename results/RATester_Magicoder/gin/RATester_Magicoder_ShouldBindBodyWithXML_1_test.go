package gin

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestShouldBindBodyWithXML_1(t *testing.T) {
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
			name:    "Valid XML",
			body:    `<testStruct><field1>value1</field1><field2>123</field2></testStruct>`,
			wantErr: false,
		},
		{
			name:    "Invalid XML",
			body:    `<testStruct><field1>value1</field1><field2>abc</field2></testStruct>`,
			wantErr: true,
		},
		{
			name:    "Empty XML",
			body:    `<testStruct></testStruct>`,
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

			var obj testStruct
			err = ctx.ShouldBindBodyWithXML(&obj)

			if tc.wantErr && err == nil {
				t.Errorf("Expected error, but got nil")
			}

			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
