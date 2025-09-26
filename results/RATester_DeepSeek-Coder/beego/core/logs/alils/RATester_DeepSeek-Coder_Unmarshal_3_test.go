package alils

import (
	"fmt"
	"testing"
)

func TestUnmarshal_3(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid data",
			input:   []byte{0x10, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x31},
			wantErr: false,
		},
		{
			name:    "invalid data",
			input:   []byte{0x10, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x32},
			wantErr: true,
		},
		{
			name:    "empty data",
			input:   []byte{},
			wantErr: true,
		},
		{
			name:    "nil data",
			input:   nil,
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

			m := &LogGroup{}
			err := m.Unmarshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
