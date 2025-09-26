package udp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestGetContent_1(t *testing.T) {
	type args struct {
		m *msg.UDPPacket
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				m: &msg.UDPPacket{
					Content: "dGVzdA==", // base64 encoded "test"
				},
			},
			wantBuf: []byte("test"),
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				m: &msg.UDPPacket{
					Content: "invalid", // invalid base64
				},
			},
			wantBuf: nil,
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

			gotBuf, err := GetContent(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("GetContent() = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}
