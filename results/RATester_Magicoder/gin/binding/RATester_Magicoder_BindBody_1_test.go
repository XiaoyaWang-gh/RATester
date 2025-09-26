package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_1(t *testing.T) {
	b := protobufBinding{}

	type TestStruct struct {
		Field1 string
		Field2 int32
	}

	testCases := []struct {
		name    string
		body    []byte
		obj     any
		wantErr bool
	}{
		{
			name: "Success",
			body: []byte{0x0A, 0x05, 0x46, 0x69, 0x65, 0x6C, 0x64, 0x31, 0x10, 0x00},
			obj: &TestStruct{
				Field1: "Field1",
				Field2: 0,
			},
			wantErr: false,
		},
		{
			name:    "Fail",
			body:    []byte{0x0A, 0x05, 0x46, 0x69, 0x65, 0x6C, 0x64, 0x31, 0x10, 0x00},
			obj:     "not a proto message",
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

			err := b.BindBody(tc.body, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindBody() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
