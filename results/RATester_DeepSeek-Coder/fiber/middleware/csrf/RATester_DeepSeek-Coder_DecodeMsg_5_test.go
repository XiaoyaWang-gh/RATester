package csrf

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestDecodeMsg_5(t *testing.T) {
	type testCase struct {
		name    string
		input   []byte
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "valid map",
			input:   []byte{0x81, 0xa1, 0x61, 0x1, 0xa1, 0x62, 0x2},
			wantErr: false,
		},
		{
			name:    "invalid map",
			input:   []byte{0x81, 0xa1, 0x61, 0x1, 0xa1, 0x62, 0x81, 0x63, 0x3},
			wantErr: true,
		},
		{
			name:    "empty map",
			input:   []byte{0x80},
			wantErr: false,
		},
		{
			name:    "invalid type",
			input:   []byte{0x90},
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

			r := msgp.NewReader(bytes.NewReader(tt.input))
			sm := &storageManager{}
			err := sm.DecodeMsg(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
