package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_4(t *testing.T) {
	type testStruct struct {
		Field string `xml:"field"`
	}

	testCases := []struct {
		name    string
		body    []byte
		obj     any
		wantErr bool
	}{
		{
			name:    "valid xml",
			body:    []byte(`<testStruct><field>test</field></testStruct>`),
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name:    "invalid xml",
			body:    []byte(`<testStruct><field>test</field></testStruct>`),
			obj:     "not a pointer",
			wantErr: true,
		},
		{
			name:    "empty body",
			body:    []byte{},
			obj:     &testStruct{},
			wantErr: false,
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
			}
		})
	}
}
