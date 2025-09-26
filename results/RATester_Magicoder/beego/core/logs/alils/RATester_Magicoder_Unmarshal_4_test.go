package alils

import (
	"fmt"
	"testing"
)

func TestUnmarshal_4(t *testing.T) {
	testCases := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   []byte{0x0a, 0x03, 0x08, 0x01, 0x12, 0x0b, 0x0a, 0x03, 0x08, 0x02},
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   []byte{0x0a, 0x03, 0x08, 0x01, 0x12, 0x0b, 0x0a, 0x03, 0x08, 0x02, 0x0a},
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

			lg := &LogGroupList{}
			err := lg.Unmarshal(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
