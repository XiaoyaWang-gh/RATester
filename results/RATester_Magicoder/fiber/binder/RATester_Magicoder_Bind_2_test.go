package binder

import (
	"fmt"
	"testing"
)

func TestBind_2(t *testing.T) {
	type testStruct struct {
		Field string `xml:"field"`
	}

	testCases := []struct {
		name    string
		body    []byte
		out     any
		wantErr bool
	}{
		{
			name: "valid xml",
			body: []byte(`<testStruct><field>value</field></testStruct>`),
			out:  &testStruct{},
		},
		{
			name:    "invalid xml",
			body:    []byte(`<testStruct><field>value</field>`),
			out:     &testStruct{},
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

			b := &xmlBinding{}
			err := b.Bind(tc.body, tc.out)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
