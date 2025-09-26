package binding

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestDecodeMsgPack_1(t *testing.T) {
	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name    string
		input   io.Reader
		wantErr bool
	}{
		{
			name: "valid input",
			input: bytes.NewBuffer([]byte{
				0x82, 0xa5, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0xa5, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x01,
			}),
			wantErr: false,
		},
		{
			name: "invalid input",
			input: bytes.NewBuffer([]byte{
				0x82, 0xa5, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0xa5, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x02,
			}),
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

			var obj testStruct
			err := decodeMsgPack(tc.input, &obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("decodeMsgPack() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
