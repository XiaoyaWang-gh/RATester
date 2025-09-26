package limiter

import (
	"fmt"
	"testing"
)

func TestUnmarshalMsg_1(t *testing.T) {
	tests := []struct {
		name    string
		z       *item
		bts     []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			z: &item{
				currHits: 10,
				prevHits: 20,
				exp:      30,
			},
			bts:     []byte{0x83, 0xa8, 0x63, 0x75, 0x72, 0x72, 0x48, 0x69, 0x74, 0x73, 0x82, 0x63, 0x75, 0x72, 0x72, 0x48, 0x69, 0x74, 0x73, 0x0a, 0x63, 0x70, 0x72, 0x65, 0x76, 0x48, 0x69, 0x74, 0x73, 0x0a, 0x65, 0x78, 0x70, 0x0a},
			wantErr: false,
		},
		// Add more test cases as needed
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
