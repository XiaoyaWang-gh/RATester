package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_6(t *testing.T) {
	type testStruct struct {
		Field1 string
		Field2 int
	}

	testCases := []struct {
		name    string
		body    []byte
		obj     any
		wantErr bool
	}{
		{
			name: "valid toml",
			body: []byte(`Field1 = "value1"
Field2 = 123`),
			obj: &testStruct{},
		},
		{
			name: "invalid toml",
			body: []byte(`Field1 = "value1"
Field2 = "not a number"`),
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

			b := tomlBinding{}
			err := b.BindBody(tc.body, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindBody() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
