package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_4(t *testing.T) {
	type testStruct struct {
		Field1 string `xml:"field1"`
		Field2 int    `xml:"field2"`
	}

	testCases := []struct {
		name    string
		body    []byte
		obj     any
		wantErr bool
	}{
		{
			name: "valid XML",
			body: []byte(`<testStruct><field1>value1</field1><field2>123</field2></testStruct>`),
			obj:  &testStruct{},
		},
		{
			name:    "invalid XML",
			body:    []byte(`<testStruct><field1>value1</field1><field2>abc</field2></testStruct>`),
			obj:     &testStruct{},
			wantErr: true,
		},
		{
			name: "empty XML",
			body: []byte(``),
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

			b := xmlBinding{}
			err := b.BindBody(tc.body, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindBody() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
