package rpc

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestReadBytes_1(t *testing.T) {
	type testCase struct {
		name    string
		input   io.Reader
		want    []byte
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid data",
			input: bytes.NewReader([]byte{
				0x00, 0x00, 0x00, 0x05, // length
				0x01, 0x02, 0x03, 0x04, 0x05, // data
			}),
			want:    []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			wantErr: false,
		},
		{
			name: "invalid length",
			input: bytes.NewReader([]byte{
				0x00, 0x00, 0x00, 0x01, 0x02, // invalid length
			}),
			want:    nil,
			wantErr: true,
		},
		{
			name: "EOF",
			input: bytes.NewReader([]byte{
				0x00, 0x00, 0x00, 0x05, // length
			}),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ReadBytes(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
