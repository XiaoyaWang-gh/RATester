package idempotency

import (
	"fmt"
	"testing"
)

func TestUnmarshalMsg_2(t *testing.T) {
	type args struct {
		bts []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestUnmarshalMsg",
			args: args{
				bts: []byte{0x83, 0xa2, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x81, 0x61, 0x81, 0x41, 0x81, 0x62, 0x81, 0x42, 0x81, 0x63, 0x81, 0x43},
			},
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

			z := &response{}
			_, err := z.UnmarshalMsg(tt.args.bts)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
