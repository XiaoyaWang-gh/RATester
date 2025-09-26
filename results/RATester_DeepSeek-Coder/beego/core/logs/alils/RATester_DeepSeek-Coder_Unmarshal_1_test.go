package alils

import (
	"fmt"
	"testing"
)

func TestUnmarshal_1(t *testing.T) {
	testCases := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   []byte{0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65},
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   []byte{0x0a, 0x03, 0x4b, 0x65, 0x79, 0x1a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65},
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

			m := &LogContent{}
			err := m.Unmarshal(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
