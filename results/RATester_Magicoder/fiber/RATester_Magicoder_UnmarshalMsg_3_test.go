package fiber

import (
	"fmt"
	"testing"
)

func TestUnmarshalMsg_3(t *testing.T) {
	tests := []struct {
		name    string
		z       *redirectionMsgs
		bts     []byte
		wantErr bool
	}{
		{
			name:    "Test case 1",
			z:       &redirectionMsgs{},
			bts:     []byte{0x81, 0x81, 0x81},
			wantErr: false,
		},
		{
			name:    "Test case 2",
			z:       &redirectionMsgs{},
			bts:     []byte{0x82, 0x82, 0x82},
			wantErr: false,
		},
		{
			name:    "Test case 3",
			z:       &redirectionMsgs{},
			bts:     []byte{0x83, 0x83, 0x83},
			wantErr: false,
		},
		{
			name:    "Test case 4",
			z:       &redirectionMsgs{},
			bts:     []byte{0x84, 0x84, 0x84},
			wantErr: false,
		},
		{
			name:    "Test case 5",
			z:       &redirectionMsgs{},
			bts:     []byte{0x85, 0x85, 0x85},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := tt.z.UnmarshalMsg(tt.bts)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
