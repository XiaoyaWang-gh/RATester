package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_3(t *testing.T) {
	b := msgpackBinding{}

	type testStruct struct {
		Field1 string `msgpack:"field1"`
		Field2 int    `msgpack:"field2"`
	}

	testCases := []struct {
		name    string
		body    []byte
		obj     any
		wantErr bool
	}{
		{
			name: "valid body",
			body: []byte{0x82, 0xa6, 'f', 'i', 'e', 'l', 'd', '1', 0xa6, 'f', 'i', 'e', 'l', 'd', '2', 0x01},
			obj:  &testStruct{},
		},
		{
			name:    "invalid body",
			body:    []byte{0x82, 0xa6, 'f', 'i', 'e', 'l', 'd', '1', 0xa6, 'f', 'i', 'e', 'l', 'd', '2', 0x01, 0x02},
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

			err := b.BindBody(tc.body, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindBody() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
