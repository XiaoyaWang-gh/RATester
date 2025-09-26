package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_7(t *testing.T) {
	type testStruct struct {
		Field1 string `yaml:"field1"`
		Field2 int    `yaml:"field2"`
	}

	testCases := []struct {
		name    string
		body    string
		obj     any
		wantErr bool
	}{
		{
			name: "valid yaml",
			body: "field1: value1\nfield2: 123",
			obj:  &testStruct{},
		},
		{
			name:    "invalid yaml",
			body:    "field1: value1\nfield2: value2",
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

			yb := yamlBinding{}
			err := yb.BindBody([]byte(tc.body), tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindBody() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
