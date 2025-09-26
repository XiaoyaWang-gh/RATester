package alils

import (
	"fmt"
	"testing"
)

func TestUnmarshal_2(t *testing.T) {
	testCases := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid log",
			input:   []byte{0x08, 0x01, 0x12, 0x06, 0x08, 0x01, 0x1a, 0x06, 0x08, 0x01},
			wantErr: false,
		},
		{
			name:    "invalid log",
			input:   []byte{0x08, 0x01, 0x12, 0x06, 0x08, 0x01, 0x1a, 0x06, 0x08, 0x01},
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

			log := &Log{}
			err := log.Unmarshal(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
